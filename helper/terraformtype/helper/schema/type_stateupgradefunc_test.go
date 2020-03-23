package schema

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestIsFuncTypeStateUpgradeFunc(t *testing.T) {
	errorIdent := &ast.Ident{
		Name: "error",
	}
	interfaceAst := &ast.InterfaceType{
		Methods: &ast.FieldList{},
	}
	stringIdent := &ast.Ident{
		Name: types.Typ[types.String].String(),
	}
	mapStringInterfaceAst := &ast.MapType{
		Key:   stringIdent,
		Value: interfaceAst,
	}
	typesInfo := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			errorIdent: types.TypeAndValue{
				Type: types.NewNamed(types.NewTypeName(token.NoPos, nil, "error", nil), nil, nil),
			},
			interfaceAst: types.TypeAndValue{
				Type: types.NewInterfaceType(nil, nil),
			},
			mapStringInterfaceAst: types.TypeAndValue{
				Type: types.NewMap(types.Typ[types.String], types.NewInterfaceType(nil, nil)),
			},
			stringIdent: types.TypeAndValue{
				Type: types.Typ[types.String],
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
			Name: "func(map[string]interface{}, interface{}) (map[string]interface{}, error)",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: mapStringInterfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: mapStringInterfaceAst,
							},
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
			Name: "func(map[string]interface{}, interface{}) (string, error)",
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: mapStringInterfaceAst,
							},
							{
								Type: interfaceAst,
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: mapStringInterfaceAst,
							},
							{
								Type: stringIdent,
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
			got := IsFuncTypeStateUpgradeFunc(testCase.Node, testCase.Info)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}

func TestIsTypeStateUpgradeFunc(t *testing.T) {
	testCases := []struct {
		Name      string
		InputType types.Type
		Expected  bool
	}{
		{
			Name: fmt.Sprintf("%s.%s", PackagePath, TypeNameStateUpgradeFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameStateUpgradeFunc, nil),
				nil,
				nil,
			),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("*%s.%s", PackagePath, TypeNameStateUpgradeFunc),
			InputType: types.NewPointer(types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameStateUpgradeFunc, nil),
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
			Name: fmt.Sprintf("%s.%s", "incorrect/path", TypeNameStateUpgradeFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage("incorrect/path", "path"), TypeNameStateUpgradeFunc, nil),
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
			got := IsTypeStateUpgradeFunc(testCase.InputType)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
