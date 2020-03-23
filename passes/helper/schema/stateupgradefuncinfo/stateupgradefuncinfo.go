package stateupgradefuncinfo

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "stateupgradefuncinfo",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/schema StateUpgradeFunc declarations for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*schema.StateUpgradeFuncInfo{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	var result []*schema.StateUpgradeFuncInfo

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if !schema.IsFuncTypeStateUpgradeFunc(n, pass.TypesInfo) {
			return
		}

		result = append(result, schema.NewStateUpgradeFuncInfo(n, pass.TypesInfo))
	})

	return result, nil
}
