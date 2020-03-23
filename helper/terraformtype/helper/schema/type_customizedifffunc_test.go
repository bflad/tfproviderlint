package schema

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestIsFuncTypeCustomizeDiffFunc(t *testing.T) {
	boolIdent := &ast.Ident{
		Name: types.Typ[types.Bool].String(),
	}
	errorIdent := &ast.Ident{
		Name: "error",
	}
	interfaceAst := &ast.InterfaceType{
		Methods: &ast.FieldList{},
	}
	packageIdent := &ast.Ident{
		Name: PackageName,
	}
	packageType := types.NewPackage(PackagePath, PackageName)
	packageNameType := types.NewPkgName(token.NoPos, packageType, packageType.Name(), packageType)
	packageFuncIdent := &ast.Ident{
		Name: TypeNameResourceDiff,
	}
	packageSelectorExpr := &ast.SelectorExpr{
		Sel: packageFuncIdent,
		X:   packageIdent,
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
			packageSelectorExpr: types.TypeAndValue{
				Type: types.NewNamed(types.NewTypeName(token.NoPos, packageType, TypeNameResourceDiff, nil), nil, nil),
			},
		},
		Uses: map[*ast.Ident]types.Object{
			packageIdent: packageNameType,
		},
	}

	testCases := []struct {
		Name     string
		Node     ast.Node
		Info     *types.Info
		Expected bool
	}{
		{
			Name: fmt.Sprintf("func(*%s.%s, interface{}) bool", PackagePath, TypeNameResourceDiff),
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: packageSelectorExpr,
								},
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
			Name: fmt.Sprintf("func(*%s.%s, interface{}) error", PackagePath, TypeNameResourceDiff),
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: packageSelectorExpr,
								},
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := IsFuncTypeCustomizeDiffFunc(testCase.Node, testCase.Info)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}

func TestIsTypeCustomizeDiffFunc(t *testing.T) {
	testCases := []struct {
		Name      string
		InputType types.Type
		Expected  bool
	}{
		{
			Name: fmt.Sprintf("%s.%s", PackagePath, TypeNameCustomizeDiffFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameCustomizeDiffFunc, nil),
				nil,
				nil,
			),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("*%s.%s", PackagePath, TypeNameCustomizeDiffFunc),
			InputType: types.NewPointer(types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameCustomizeDiffFunc, nil),
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
			Name: fmt.Sprintf("%s.%s", "incorrect/path", TypeNameCustomizeDiffFunc),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage("incorrect/path", "path"), TypeNameCustomizeDiffFunc, nil),
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
			got := IsTypeCustomizeDiffFunc(testCase.InputType)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
