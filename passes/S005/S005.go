// Package S005 defines an Analyzer that checks for
// Schema with Computed enabled and Default configured
package S005

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaschema"
)

const Doc = `check for Schema with Computed enabled and Default configured

The S005 analyzer reports cases of schemas which enables Computed
and configures Default, which will fail provider schema validation.`

const analyzerName = "S005"

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

		computed := terraformtype.HelperSchemaTypeSchemaComputed(schema)

		if computed == nil || !*computed {
			continue
		}

		schemaDefault := terraformtype.HelperSchemaTypeSchemaDefault(schema)

		if schemaDefault == nil {
			continue
		}

		switch t := schema.Type.(type) {
		default:
			pass.Reportf(schema.Lbrace, "%s: schema should not enable Computed and configure Default", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema should not enable Computed and configure Default", analyzerName)
		}
	}

	return nil, nil
}
