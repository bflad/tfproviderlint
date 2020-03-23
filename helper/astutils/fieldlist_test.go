package astutils

import (
	"go/ast"
	"testing"
)

func TestHasFieldListLength(t *testing.T) {
	testCases := []struct {
		Name           string
		FieldList      *ast.FieldList
		ExpectedLength int
		Expected       bool
	}{
		{
			Name:           "nil with length",
			FieldList:      nil,
			ExpectedLength: 1,
			Expected:       false,
		},
		{
			Name:           "nil without length",
			FieldList:      nil,
			ExpectedLength: 0,
			Expected:       true,
		},
		{
			Name: "correct length",
			FieldList: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.Ident{
							Name: "bool",
						},
					},
					{
						Type: &ast.Ident{
							Name: "bool",
						},
					},
				},
			},
			ExpectedLength: 2,
			Expected:       true,
		},
		{
			Name: "incorrect length",
			FieldList: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.Ident{
							Name: "bool",
						},
					},
					{
						Type: &ast.Ident{
							Name: "bool",
						},
					},
				},
			},
			ExpectedLength: 1,
			Expected:       false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := HasFieldListLength(testCase.FieldList, testCase.ExpectedLength)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
