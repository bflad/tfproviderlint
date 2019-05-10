// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package unsafeptr defines an Analyzer that checks for invalid
// conversions of uintptr to unsafe.Pointer.
package acctest

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check for invalid conversions of uintptr to unsafe.Pointer

The unsafeptr analyzer reports likely incorrect uses of unsafe.Pointer
to convert integers to pointers. A conversion from uintptr to
unsafe.Pointer is invalid if it implies that there is a uintptr-typed
word in memory that holds a pointer value, because that word will be
invisible to stack copying and to the garbage collector.`

var Analyzer = &analysis.Analyzer{
	Name:     "acctest",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		x := n.(*ast.CompositeLit)
		var pos token.Pos

		switch v := x.Type.(type) {
		default:
			return
		case *ast.SelectorExpr:
			if v.Sel.Name != "TestCase" {
				return
			}
			pos = v.Sel.Pos()
		}

		for _, elt := range x.Elts {
			switch v := elt.(type) {
			default:
				return
			case *ast.KeyValueExpr:
				if v.Key.(*ast.Ident).Name == "CheckDestroy" {
					return
				}
			}
		}

		pass.Reportf(pos, "missing CheckDestroy")
	})
	return nil, nil
}
