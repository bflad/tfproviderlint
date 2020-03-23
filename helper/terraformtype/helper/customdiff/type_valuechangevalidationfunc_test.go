package customdiff

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestIsFuncTypeValueChangeValidationFunc(t *testing.T) {
	boolIdent := &ast.Ident{
		Name: types.Typ[types.Bool].String(),
	}
	errorIdent := &ast.Ident{
		Name: "error",
	}
	interfaceAst := &ast.InterfaceType{
		Methods: &ast.FieldList{},
	}
	typesInfo := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			boolIdent: types.TypeAndValue{
				Type: types.Typ[types.Bool],
			},
			errorIdent: types.TypeAndValue{
				Type: types.NewNamed(types.NewTypeName(token.NoPos, nil, "error", nil), nil, nil),
			},
			interfaceAst: types.TypeAndValue{
				Type: types.NewInterfaceType(nil, nil),
			},
		},
	}

	testCases := []struct {
		Name     string
		Node     ast.Node
		Info     *types.Info
		Expected bool
	}{
		{
			Name: "func(interface{}, interface{}, interface{}) bool",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: boolIdent,
							},
						},
					},
				},
			},
			Info:     typesInfo,
			Expected: false,
		},
		{
			Name: "func(interface{}, interface{}, interface{}) error",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: errorIdent,
							},
						},
					},
				},
			},
			Info:     typesInfo,
			Expected: true,
		},
		{
			Name: "func(interface{}, interface{}) bool",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: boolIdent,
							},
						},
					},
				},
			},
			Info:     typesInfo,
			Expected: false,
		},
		{
			Name: "func(interface{}, interface{}) error",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: interfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: errorIdent,
							},
						},
					},
				},
			},
			Info:     typesInfo,
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := IsFuncTypeValueChangeValidationFunc(testCase.Node, testCase.Info)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}

func TestIsTypeValueChangeValidationFunc(t *testing.T) {
	testCases := []struct {
		Name      string
		InputType types.Type
		Expected  bool
	}{
		{
			Name: fmt.Sprintf("%s.%s", PackagePath, TypeNameValueChangeValidationFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameValueChangeValidationFunc, nil),
				nil,
				nil,
			),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("*%s.%s", PackagePath, TypeNameValueChangeValidationFunc),
			InputType: types.NewPointer(types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameValueChangeValidationFunc, nil),
				nil,
				nil,
			)),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("%s.%s", PackagePath, "Not"),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), "Not", nil),
				nil,
				nil,
			),
			Expected: false,
		},
		{
			Name: fmt.Sprintf("%s.%s", "incorrect/path", TypeNameValueChangeValidationFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage("incorrect/path", "path"), TypeNameValueChangeValidationFunc, nil),
				nil,
				nil,
			),
			Expected: false,
		},
		{
			Name:      "string",
			InputType: types.Typ[types.String],
			Expected:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := IsTypeValueChangeValidationFunc(testCase.InputType)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
