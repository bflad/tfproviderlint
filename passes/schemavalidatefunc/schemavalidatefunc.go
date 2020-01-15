package schemavalidatefunc

import (
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "schemavalidatefunc",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/schema SchemaValidateFunc declarations for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*ast.FuncDecl{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	var result []*ast.FuncDecl

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		funcDecl := n.(*ast.FuncDecl)

		params := funcDecl.Type.Params.List

		if len(params) != 2 {
			return
		}

		if !astutils.IsFunctionParameterTypeInterface(params[0].Type) {
			return
		}

		if !astutils.IsFunctionParameterTypeString(params[1].Type) {
			return
		}

		results := funcDecl.Type.Results.List

		if len(results) != 2 {
			return
		}

		if !astutils.IsFunctionParameterTypeArrayString(results[0].Type) {
			return
		}

		if !astutils.IsFunctionParameterTypeArrayError(results[1].Type) {
			return
		}

		result = append(result, funcDecl)
	})

	return result, nil
}
