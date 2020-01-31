// Package S024 defines an Analyzer that checks for
// extraneous usage of ForceNew in data source schema attributes
package S024

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaresource"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Schema that should omit ForceNew in data source schema attributes

The S024 analyzer reports usage of ForceNew in data source schema attributes,
which is unnecessary.`

const analyzerName = "S024"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemaresource.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resources := pass.ResultOf[schemaresource.Analyzer].([]*terraformtype.HelperSchemaResourceInfo)
	for _, resource := range resources {
		if ignorer.ShouldIgnore(analyzerName, resource.AstCompositeLit) {
			continue
		}

		if !resource.IsDataSource() {
			continue
		}

		var schemas []*terraformtype.HelperSchemaSchemaInfo

		ast.Inspect(resource.AstCompositeLit, func(n ast.Node) bool {
			compositeLit, ok := n.(*ast.CompositeLit)

			if !ok {
				return true
			}

			if terraformtype.IsMapStringHelperSchemaTypeSchema(compositeLit, pass.TypesInfo) {
				for _, schema := range terraformtype.GetSchemaMapSchemas(compositeLit) {
					schemas = append(schemas, terraformtype.NewHelperSchemaSchemaInfo(schema, pass.TypesInfo))
				}
			} else if terraformtype.IsHelperSchemaTypeSchema(pass.TypesInfo.TypeOf(compositeLit.Type)) {
				schemas = append(schemas, terraformtype.NewHelperSchemaSchemaInfo(compositeLit, pass.TypesInfo))
			}

			return true
		})

		for _, schema := range schemas {
			if !schema.DeclaresField(terraformtype.SchemaFieldForceNew) {
				continue
			}

			pass.Reportf(schema.Fields[terraformtype.SchemaFieldForceNew].Pos(), "%s: ForceNew is extraneous in data source schema attributes", analyzerName)
		}
	}

	return nil, nil
}
