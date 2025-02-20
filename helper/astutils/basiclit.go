package astutils

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
)

// ExprBoolConstValue fetches a bool value from the Expr
// If the Expr is not a bool constant, returns nil.
func ExprBoolConstValue(e ast.Expr, info *types.Info) *bool {
	t, ok := info.Types[e]
	if !ok {
		return nil
	}
	typ := t.Type
	if tt, ok := typ.(*types.Named); ok {
		typ = tt.Underlying()
	}
	if bt, ok := typ.(*types.Basic); !ok || bt.Kind() != types.Bool {
		return nil
	}
	// Constant value is guaranteed non-nil
	if t.Value == nil {
		return nil
	}
	boolValue, err := strconv.ParseBool(t.Value.ExactString())
	if err != nil {
		return nil
	}
	return &boolValue
}

// ExprIntConstValue fetches an int value from the Expr
// If the Expr is not an int constant, returns nil.
func ExprIntConstValue(e ast.Expr, info *types.Info) *int {
	t, ok := info.Types[e]
	if !ok {
		return nil
	}
	typ := t.Type
	if tt, ok := typ.(*types.Named); ok {
		typ = tt.Underlying()
	}
	if bt, ok := typ.(*types.Basic); !ok || bt.Kind() != types.Int {
		return nil
	}
	// Constant value is guaranteed non-nil
	if t.Value == nil {
		return nil
	}
	intValue, err := strconv.Atoi(t.Value.ExactString())
	if err != nil {
		return nil
	}
	return &intValue
}

// ExprStringConstValue fetches a string value from the Expr
// If the Expr is not a string constant, returns nil.
func ExprStringConstValue(e ast.Expr, info *types.Info) *string {
	t, ok := info.Types[e]
	if !ok {
		return nil
	}
	typ := t.Type
	if tt, ok := typ.(*types.Named); ok {
		typ = tt.Underlying()
	}
	if bt, ok := typ.(*types.Basic); !ok || bt.Kind() != types.String {
		return nil
	}
	// Constant value is guaranteed non-nil
	if t.Value == nil {
		return nil
	}
	stringValue, _ := strconv.Unquote(t.Value.ExactString()) // can assume well-formed Go
	return &stringValue
}

// ExprBoolValue fetches a bool value from the Expr
// If the Expr cannot parse as a bool, returns nil.
func ExprBoolValue(e ast.Expr) *bool {
	switch v := e.(type) {
	case *ast.Ident:
		stringValue := v.Name
		boolValue, err := strconv.ParseBool(stringValue)

		if err != nil {
			return nil
		}

		return &boolValue
	}
	return nil
}

// ExprIntValue fetches an int value from the Expr
// If the Expr cannot parse as an int, returns nil.
func ExprIntValue(e ast.Expr) *int {
	switch v := e.(type) {
	case *ast.BasicLit:
		intValue, err := strconv.Atoi(v.Value)

		if err != nil {
			return nil
		}

		return &intValue
	}

	return nil
}

// ExprStringValue fetches a string value from the Expr
// If the Expr is not BasicLit, returns an empty string.
func ExprStringValue(e ast.Expr) *string {
	switch v := e.(type) {
	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return nil
		}
		stringValue, _ := strconv.Unquote(v.Value) // can assume well-formed Go
		return &stringValue
	}

	return nil
}

// ExprValue fetches a pointer to the Expr
// If the Expr is nil, returns nil
func ExprValue(e ast.Expr) *ast.Expr {
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
