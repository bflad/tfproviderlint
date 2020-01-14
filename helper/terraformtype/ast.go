package terraformtype

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"
)

func astBoolValue(e ast.Expr) *bool {
	switch v := e.(type) {
	case *ast.Ident:
		stringValue := v.Name
		boolValue, err := strconv.ParseBool(stringValue)

		if err != nil {
			return nil
		}

		return &boolValue
	default:
		return nil
	}
}

func astCompositeLitField(cl *ast.CompositeLit, fieldName string) *ast.KeyValueExpr {
	for _, elt := range cl.Elts {
		switch e := elt.(type) {
		case *ast.KeyValueExpr:
			if e.Key.(*ast.Ident).Name != fieldName {
				continue
			}

			return e
		}
	}

	return nil
}

func astCompositeLitFields(cl *ast.CompositeLit) map[string]*ast.KeyValueExpr {
	result := make(map[string]*ast.KeyValueExpr, len(cl.Elts))

	for _, elt := range cl.Elts {
		switch e := elt.(type) {
		case *ast.KeyValueExpr:
			result[e.Key.(*ast.Ident).Name] = e
		}
	}

	return result
}

func astCompositeLitFieldBoolValue(cl *ast.CompositeLit, fieldName string) *bool {
	kvExpr := astCompositeLitField(cl, fieldName)

	if kvExpr == nil {
		return nil
	}

	return astBoolValue(kvExpr.Value)
}

func astCompositeLitFieldExprValue(cl *ast.CompositeLit, fieldName string) *ast.Expr {
	kvExpr := astCompositeLitField(cl, fieldName)

	if kvExpr == nil {
		return nil
	}

	return astExprValue(kvExpr.Value)
}

func astCompositeLitFieldIntValue(cl *ast.CompositeLit, fieldName string) *int {
	kvExpr := astCompositeLitField(cl, fieldName)

	if kvExpr == nil {
		return nil
	}

	return astIntValue(kvExpr.Value)
}

func astCompositeLitContainsAnyField(cl *ast.CompositeLit, fieldNames ...string) bool {
	for _, elt := range cl.Elts {
		switch e := elt.(type) {
		case *ast.KeyValueExpr:
			name := e.Key.(*ast.Ident).Name

			for _, field := range fieldNames {
				if name == field {
					return true
				}
			}
		}
	}

	return false
}

func astExprValue(e ast.Expr) *ast.Expr {
	switch v := e.(type) {
	case *ast.Ident:
		if v.Name == "nil" {
			return nil
		}

		return &e
	default:
		return &e
	}
}

func astIntValue(e ast.Expr) *int {
	switch v := e.(type) {
	case *ast.BasicLit:
		stringValue := strings.Trim(v.Value, `"`)
		intValue, err := strconv.Atoi(stringValue)

		if err != nil {
			return nil
		}

		return &intValue
	default:
		return nil
	}
}

// isPackageReceiverMethod returns true if the package suffix (for vendoring), receiver name, and method name match
func isPackageReceiverMethod(e ast.Expr, info *types.Info, packageSuffix string, receiverName, methodName string) bool {
	switch e := e.(type) {
	case *ast.SelectorExpr:
		if e.Sel.Name != methodName {
			return false
		}

		switch x := e.X.(type) {
		case *ast.Ident:
			if x.Obj == nil {
				return false
			}

			switch decl := x.Obj.Decl.(type) {
			case *ast.Field:
				switch t := decl.Type.(type) {
				case *ast.StarExpr:
					return isPackageType(info.TypeOf(t.X), packageSuffix, receiverName)
				case *ast.SelectorExpr:
					return isPackageType(info.TypeOf(t), packageSuffix, receiverName)
				}
			case *ast.ValueSpec:
				return isPackageType(info.TypeOf(decl.Type), packageSuffix, receiverName)
			}
		}
	}

	return false
}

// isPackageFunc returns true if the function package suffix (for vendoring) and name matches
func isPackageFunc(e ast.Expr, info *types.Info, packageSuffix string, funcName string) bool {
	switch e := e.(type) {
	case *ast.SelectorExpr:
		if e.Sel.Name != funcName {
			return false
		}

		switch x := e.X.(type) {
		case *ast.Ident:
			return strings.HasSuffix(info.ObjectOf(x).(*types.PkgName).Imported().Path(), packageSuffix)
		}
	}

	return false
}

// isPackageNamedType returns if the type name matches and is from the package suffix
func isPackageNamedType(t *types.Named, packageSuffix string, typeName string) bool {
	if t.Obj().Name() != typeName {
		return false
	}

	// HasSuffix here due to vendoring
	return strings.HasSuffix(t.Obj().Pkg().Path(), packageSuffix)
}

// isPackageType returns true if the type name can be matched and is from the package suffix
func isPackageType(t types.Type, packageSuffix string, typeName string) bool {
	switch t := t.(type) {
	case *types.Named:
		return isPackageNamedType(t, packageSuffix, typeName)
	}

	return false
}
