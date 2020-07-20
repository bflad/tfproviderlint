package XS002

import (
	"go/ast"
	"go/token"
	"sort"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemamapcompositelit"
)

const Doc = `check for Schema that attribute names are in alphabetical order

The XS002 analyzer reports cases of schemas where attribute names
are not in alphabetical order.`

const analyzerName = "XS002"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemamapcompositelit.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemamapcompositelits := pass.ResultOf[schemamapcompositelit.Analyzer].([]*ast.CompositeLit)

	for _, smap := range schemamapcompositelits {
		if ignorer.ShouldIgnore(analyzerName, smap) {
			continue
		}

		schemaKeys := make([]string, 0 , len(schema.GetSchemaMapAttributeNames(smap)))
		schemaKeyPos := make([]token.Pos, 0 , len(schema.GetSchemaMapAttributeNames(smap)))

		for _, attributeName := range schema.GetSchemaMapAttributeNames(smap) {
			schemaKeys = append(schemaKeys, attributeName.(*ast.BasicLit).Value)
			schemaKeyPos = append(schemaKeyPos, attributeName.(*ast.BasicLit).Pos())
		}

		if !sort.StringsAreSorted(schemaKeys) {
			sortedSchemaKeys := make([]string, 0, len(schemaKeys))
			for _, k := range schemaKeys {
				sortedSchemaKeys = append(sortedSchemaKeys, k)
			}
			sort.Strings(sortedSchemaKeys)
			// find occurrences out of alphabetical order
			var pos []token.Pos
			for i, v := range schemaKeys {
				if v != sortedSchemaKeys[i] {
					pos = append(pos, schemaKeyPos[i])
				}
			}
			for _, p := range pos{
				pass.Reportf(p, "%s: schema attribute name should be in alphabetical order", analyzerName)
			}

		}
	}

	return nil, nil
}