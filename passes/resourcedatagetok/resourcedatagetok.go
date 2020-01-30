package resourcedatagetok

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "resourcedatagetok",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/schema.ResourceData.GetOk() calls for later passes",
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

		if !terraformtype.IsHelperSchemaReceiverMethod(callExpr.Fun, pass.TypesInfo, terraformtype.TypeNameResourceData, "GetOk") {
			return
		}

		result = append(result, callExpr)
	})

	return result, nil
}
