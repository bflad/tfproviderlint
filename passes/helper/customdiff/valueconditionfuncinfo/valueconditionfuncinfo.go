package valueconditionfuncinfo

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/customdiff"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "valueconditionfuncinfo",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/customdiff ValueConditionFunc declarations for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*customdiff.ValueConditionFuncInfo{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	var result []*customdiff.ValueConditionFuncInfo

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if !customdiff.IsFuncTypeValueConditionFunc(n, pass.TypesInfo) {
			return
		}

		result = append(result, customdiff.NewValueConditionFuncInfo(n, pass.TypesInfo))
	})

	return result, nil
}
