// Package S001 defines an Analyzer that checks for
// Schema of TypeList or TypeSet missing Elem
package S001

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaschema"
)

const Doc = `check for Schema of TypeList or TypeSet missing Elem

The S001 analyzer reports cases of TypeList or TypeSet schemas missing Elem,
which will fail schema validation.`

const analyzerName = "S001"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemaschema.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemas := pass.ResultOf[schemaschema.Analyzer].([]*ast.CompositeLit)
	for _, schema := range schemas {
		if ignorer.ShouldIgnore(analyzerName, schema) {
			continue
		}

		if terraformtype.HelperSchemaTypeSchemaContainsFields(schema, terraformtype.SchemaFieldElem) {
			continue
		}

		if !terraformtype.HelperSchemaTypeSchemaContainsTypes(schema, pass.TypesInfo, terraformtype.SchemaValueTypeList, terraformtype.SchemaValueTypeSet) {
			continue
		}

		switch t := schema.Type.(type) {
		default:
			pass.Reportf(schema.Lbrace, "%s: schema of TypeList or TypeSet should include Elem", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema of TypeList or TypeSet should include Elem", analyzerName)
		}
	}

	return nil, nil
}
