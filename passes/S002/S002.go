// Package S002 defines an Analyzer that checks for
// Schema with both Required and Optional enabled
package S002

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaschema"
)

const Doc = `check for Schema with both Required and Optional enabled

The S002 analyzer reports cases of schemas which enables both Required
and Optional, which will fail provider schema validation.`

const analyzerName = "S002"

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

		optional := terraformtype.HelperSchemaTypeSchemaOptional(schema)

		if optional == nil || !*optional {
			continue
		}

		required := terraformtype.HelperSchemaTypeSchemaRequired(schema)

		if required == nil || !*required {
			continue
		}

		switch t := schema.Type.(type) {
		default:
			pass.Reportf(schema.Lbrace, "%s: schema should not enable Required and Optional", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema should not enable Required and Optional", analyzerName)
		}
	}

	return nil, nil
}
