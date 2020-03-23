package valuechangeconditionfuncinfo

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/customdiff"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "valuechangeconditionfuncinfo",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/customdiff ValueChangeConditionFunc declarations for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*customdiff.ValueChangeConditionFuncInfo{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	var result []*customdiff.ValueChangeConditionFuncInfo

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if !customdiff.IsFuncTypeValueChangeConditionFunc(n, pass.TypesInfo) {
			return
		}

		result = append(result, customdiff.NewValueChangeConditionFuncInfo(n, pass.TypesInfo))
	})

	return result, nil
}
