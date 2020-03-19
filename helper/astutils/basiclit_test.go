package astutils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestExprBoolValue(t *testing.T) {
	cases := []struct {
		desc       string
		assignExpr string
		expect     *bool
	}{
		{
			"bool literal",
			`_ = true`,
			newPtr(true).(*bool),
		},
		{
			"not bool literal",
			`_ = "true"`,
			nil,
		},
		{
			"nil",
			"_ = nil",
			nil,
		},
	}

	code := `package main

func f() {
`
	for _, c := range cases {
		code += "\t" + c.assignExpr + "\n"
	}
	code += "}"

	var fset token.FileSet
	f, _ := parser.ParseFile(&fset, "foo.go", code, 0)
	mainBody := f.Decls[0].(*ast.FuncDecl).Body.List
	for i, c := range cases {
		astmt := mainBody[i].(*ast.AssignStmt)
		actual := ExprBoolValue(astmt.Rhs[0])
		switch {
		case actual == nil && c.expect == nil:
			continue
		case actual != nil && c.expect != nil:
			if *actual == *c.expect {
				continue
			}
			t.Fatalf("%s: %v (expect: %v)", c.desc, *actual, *c.expect)
		case actual == nil && c.expect != nil:
			t.Fatalf("%s: nil (expect: %v)", c.desc, *c.expect)
		case actual != nil && c.expect == nil:
			t.Fatalf("%s: %v (expect: nil)", c.desc, *actual)
		}
	}
}

func TestExprIntValue(t *testing.T) {
	cases := []struct {
		desc       string
		assignExpr string
		expect     *int
	}{
		{
			"int literal",
			`_ = 123`,
			newPtr(123).(*int),
		},
		{
			"not int literal (string)",
			`_ = "123"`,
			nil,
		},
		{
			"not int literal (bool)",
			`_ = false`,
			nil,
		},
		{
			"nil",
			"_ = nil",
			nil,
		},
	}

	code := `package main

func f() {
`
	for _, c := range cases {
		code += "\t" + c.assignExpr + "\n"
	}
	code += "}"

	var fset token.FileSet
	f, _ := parser.ParseFile(&fset, "foo.go", code, 0)
	mainBody := f.Decls[0].(*ast.FuncDecl).Body.List
	for i, c := range cases {
		astmt := mainBody[i].(*ast.AssignStmt)
		actual := ExprIntValue(astmt.Rhs[0])
		switch {
		case actual == nil && c.expect == nil:
			continue
		case actual != nil && c.expect != nil:
			if *actual == *c.expect {
				continue
			}
			t.Fatalf("%s: %v (expect: %v)", c.desc, *actual, *c.expect)
		case actual == nil && c.expect != nil:
			t.Fatalf("%s: nil (expect: %v)", c.desc, *c.expect)
		case actual != nil && c.expect == nil:
			t.Fatalf("%s: %v (expect: nil)", c.desc, *actual)
		}
	}
}

func TestExprStringValue(t *testing.T) {
	cases := []struct {
		desc       string
		assignExpr string
		expect     *string
	}{
		{
			"string with double quotes",
			`_ = "abc"`,
			newPtr("abc").(*string),
		},
		{
			"string with backtick quotes",
			"_ = `abc`",
			newPtr("abc").(*string),
		},
		{
			"not string literal",
			"_ = 123",
			nil,
		},
		{
			"nil",
			"_ = nil",
			nil,
		},
	}

	code := `package main

func f() {
`
	for _, c := range cases {
		code += "\t" + c.assignExpr + "\n"
	}
	code += "}"

	var fset token.FileSet
	f, _ := parser.ParseFile(&fset, "foo.go", code, 0)
	mainBody := f.Decls[0].(*ast.FuncDecl).Body.List
	for i, c := range cases {
		astmt := mainBody[i].(*ast.AssignStmt)
		actual := ExprStringValue(astmt.Rhs[0])
		switch {
		case actual == nil && c.expect == nil:
			continue
		case actual != nil && c.expect != nil:
			if *actual == *c.expect {
				continue
			}
			t.Fatalf("%s: %v (expect: %v)", c.desc, *actual, *c.expect)
		case actual == nil && c.expect != nil:
			t.Fatalf("%s: nil (expect: %v)", c.desc, *c.expect)
		case actual != nil && c.expect == nil:
			t.Fatalf("%s: %v (expect: nil)", c.desc, *actual)
		}
	}
}

func newPtr(v interface{}) interface{} {
	rv := reflect.New(reflect.ValueOf(v).Type())
	rv.Elem().Set(reflect.ValueOf(v))
	return rv.Interface()
}
