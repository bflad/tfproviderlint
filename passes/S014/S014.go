// Package S014 defines an Analyzer that checks for
// Schema that within Elem, Computed, Optional, and Required
// are not configured
package S014

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemamap"
)

const Doc = `check for Schema that Elem does not contain extraneous fields

The S014 analyzer reports cases of schemas which within Elem, that
Computed, Optional, and Required are not configured, which will fail
provider schema validation.`

const analyzerName = "S014"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemamap.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemamaps := pass.ResultOf[schemamap.Analyzer].([]*ast.CompositeLit)

	for _, smap := range schemamaps {
		for _, schemaCompositeLit := range terraformtype.GetSchemaMapSchemas(smap) {
			schema := terraformtype.NewHelperSchemaSchemaInfo(schemaCompositeLit, pass.TypesInfo)

			if ignorer.ShouldIgnore(analyzerName, schema.AstCompositeLit) {
				continue
			}

			elemKvExpr := schema.Fields[terraformtype.SchemaFieldElem]

			if elemKvExpr == nil {
				continue
			}

			// search within Elem
			switch elemValue := elemKvExpr.Value.(type) {
			default:
				continue
			case *ast.UnaryExpr:
				if elemValue.Op != token.AND || !terraformtype.IsHelperSchemaTypeSchema(pass.TypesInfo.TypeOf(elemValue.X)) {
					continue
				}

				switch tElemSchema := elemValue.X.(type) {
				default:
					continue
				case *ast.CompositeLit:
					elemSchema := terraformtype.NewHelperSchemaSchemaInfo(tElemSchema, pass.TypesInfo)

					for _, field := range []string{terraformtype.SchemaFieldComputed, terraformtype.SchemaFieldOptional, terraformtype.SchemaFieldRequired} {
						if kvExpr := elemSchema.Fields[field]; kvExpr != nil {
							pass.Reportf(kvExpr.Pos(), "%s: schema within Elem should not configure Computed, Optional, or Required", analyzerName)
							break
						}
					}
				}
			}
		}
	}

	return nil, nil
}
