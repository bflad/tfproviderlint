package customizedifffuncinfo

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "customizedifffuncinfo",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/schema CustomizeDiffFunc declarations for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*schema.CustomizeDiffFuncInfo{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	var result []*schema.CustomizeDiffFuncInfo

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if !schema.IsFuncTypeCustomizeDiffFunc(n, pass.TypesInfo) {
			return
		}

		result = append(result, schema.NewCustomizeDiffFuncInfo(n, pass.TypesInfo))
	})

	return result, nil
}
