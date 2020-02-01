// Package XS001 defines an Analyzer that checks for
// Schema that Description is configured
package XS001

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemamap"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Schema that Description is configured

The XS001 analyzer reports cases of schemas where Description is not
configured, which is generally useful for providers that wish to
automatically generate documentation based on the schema information.`

const analyzerName = "XS001"

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
		for _, schemaCompositeLit := range schema.GetSchemaMapSchemas(smap) {
			schemaInfo := schema.NewSchemaInfo(schemaCompositeLit, pass.TypesInfo)

			if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
				continue
			}

			if schemaInfo.Schema.Description != "" {
				continue
			}

			switch t := schemaInfo.AstCompositeLit.Type.(type) {
			default:
				pass.Reportf(schemaInfo.AstCompositeLit.Lbrace, "%s: schema should configure Description", analyzerName)
			case *ast.SelectorExpr:
				pass.Reportf(t.Sel.Pos(), "%s: schema should configure Description", analyzerName)
			}
		}
	}

	return nil, nil
}
