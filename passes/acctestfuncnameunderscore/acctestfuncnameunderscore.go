// Package acctestfuncnameunderscore defines an Analyzer that checks for
// acceptance test names missing an underscore
package acctestfuncnameunderscore

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestfunc"
	"github.com/terraform-providers/terraform-provider-aws/linter/passes/commentignore"
)

const Doc = `check for acceptance test function names missing an underscore

The acctestfuncnameunderscore analyzer reports where an underscore is not
present in the function name, which could make per-resource testing harder to
execute in larger providers or those with overlapping resource names.`

const analyzerName = "acctestfuncnameunderscore"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		acctestfunc.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	testAccFuncs := pass.ResultOf[acctestfunc.Analyzer].([]*ast.FuncDecl)
	for _, testAccFunc := range testAccFuncs {
		if ignorer.ShouldIgnore(analyzerName, testAccFunc) {
			continue
		}

		if !strings.Contains(testAccFunc.Name.Name, "_") {
			pass.Reportf(testAccFunc.Name.NamePos, "acceptance test function name should include underscore")
		}
	}

	return nil, nil
}
