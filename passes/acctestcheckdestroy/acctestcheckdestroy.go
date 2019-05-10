// Package acctestcheckdestroy defines an Analyzer that checks for
// TestCase missing CheckDestroy
package acctestcheckdestroy

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check for TestCase missing CheckDestroy

The acctestcheckdestroy analyzer reports likely incorrect uses of TestCase
which do not define a CheckDestroy function. CheckDestroy is used to verify
that test infrastructure has been removed at the end of an acceptance test.

More information can be found at:
https://www.terraform.io/docs/extend/testing/acceptance-tests/testcase.html#checkdestroy`

var Analyzer = &analysis.Analyzer{
	Name:     "acctest",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		x := n.(*ast.CompositeLit)
		var pos token.Pos

		switch v := x.Type.(type) {
		default:
			return
		case *ast.SelectorExpr:
			if v.Sel.Name != "TestCase" {
				return
			}
			pos = v.Sel.Pos()
		}

		for _, elt := range x.Elts {
			switch v := elt.(type) {
			default:
				return
			case *ast.KeyValueExpr:
				if v.Key.(*ast.Ident).Name == "CheckDestroy" {
					return
				}
			}
		}

		pass.Reportf(pos, "missing CheckDestroy")
	})
	return nil, nil
}
