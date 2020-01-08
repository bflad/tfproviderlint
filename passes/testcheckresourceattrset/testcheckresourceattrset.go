package testcheckresourceattrset

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "testcheckresourceattrset",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/resource.TestCheckResourceAttrSet calls for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*ast.CallExpr{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	var result []*ast.CallExpr

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)

		if !terraformtype.IsHelperResourceFunc(callExpr.Fun, pass.TypesInfo, terraformtype.FuncNameTestCheckResourceAttrSet) {
			return
		}

		result = append(result, callExpr)
	})

	return result, nil
}
