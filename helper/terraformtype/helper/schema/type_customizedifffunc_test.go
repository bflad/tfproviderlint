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
	contextIdent := &ast.Ident{
		Name: "context",
	}
	contextContextIdent := &ast.Ident{
		Name: "Context",
	}
	errorIdent := &ast.Ident{
		Name: "error",
	}
	interfaceAst := &ast.InterfaceType{
		Methods: &ast.FieldList{},
	}
	packageIdentV1 := &ast.Ident{
		Name: PackageName,
	}
	packageIdentV2 := &ast.Ident{
		Name: PackageName + `2`,
	}
	contextType := types.NewPackage("context", "context")
	packageTypeV1 := types.NewPackage(PackagePathVersion(1), PackageName)
	packageTypeV2 := types.NewPackage(PackagePathVersion(2), PackageName)
	contextNameType := types.NewPkgName(token.NoPos, contextType, contextType.Name(), contextType)
	packageNameTypeV1 := types.NewPkgName(token.NoPos, packageTypeV1, packageTypeV1.Name(), packageTypeV1)
	packageNameTypeV2 := types.NewPkgName(token.NoPos, packageTypeV2, packageTypeV2.Name(), packageTypeV2)
	packageFuncIdent := &ast.Ident{
		Name: TypeNameResourceDiff,
	}
	contextContextSelectorExpr := &ast.SelectorExpr{
		Sel: contextContextIdent,
		X:   contextIdent,
	}
	packageSelectorExprV1 := &ast.SelectorExpr{
		Sel: packageFuncIdent,
		X:   packageIdentV1,
	}
	packageSelectorExprV2 := &ast.SelectorExpr{
		Sel: packageFuncIdent,
		X:   packageIdentV2,
	}
	typesInfo := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			boolIdent: {
				Type: types.Typ[types.Bool],
			},
			contextContextSelectorExpr: {
				Type: types.NewNamed(types.NewTypeName(token.NoPos, contextType, contextContextIdent.Name, nil), nil, nil),
			},
			errorIdent: {
				Type: types.NewNamed(types.NewTypeName(token.NoPos, nil, "error", nil), nil, nil),
			},
			interfaceAst: {
				Type: types.NewInterfaceType(nil, nil),
			},
			packageSelectorExprV1: {
				Type: types.NewNamed(types.NewTypeName(token.NoPos, packageTypeV1, TypeNameResourceDiff, nil), nil, nil),
			},
			packageSelectorExprV2: {
				Type: types.NewNamed(types.NewTypeName(token.NoPos, packageTypeV2, TypeNameResourceDiff, nil), nil, nil),
			},
		},
		Uses: map[*ast.Ident]types.Object{
			contextIdent:   contextNameType,
			packageIdentV1: packageNameTypeV1,
			packageIdentV2: packageNameTypeV2,
		},
	}

	testCases := []struct {
		Name     string
		Node     ast.Node
		Info     *types.Info
		Expected bool
	}{
		{
			Name: fmt.Sprintf("func(*%s.%s, interface{}) bool", PackagePathVersion(1), TypeNameResourceDiff),
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: packageSelectorExprV1,
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
			Name: fmt.Sprintf("func(*%s.%s, interface{}) error", PackagePathVersion(1), TypeNameResourceDiff),
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: packageSelectorExprV1,
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
		{
			Name: fmt.Sprintf("func(context.Context, *%s.%s, interface{}) error", PackagePathVersion(2), TypeNameResourceDiff),
			Node: &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: contextContextSelectorExpr,
							},
							{
								Type: &ast.StarExpr{
									X: packageSelectorExprV2,
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
