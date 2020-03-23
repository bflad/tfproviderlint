package schema

import (
	"fmt"
	"go/token"
	"go/types"
	"testing"
)

func TestIsTypeResourceDiff(t *testing.T) {
	testCases := []struct {
		Name      string
		InputType types.Type
		Expected  bool
	}{
		{
			Name: fmt.Sprintf("%s.%s", PackagePath, TypeNameResourceDiff),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameResourceDiff, nil),
				nil,
				nil,
			),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("*%s.%s", PackagePath, TypeNameResourceDiff),
			InputType: types.NewPointer(types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePath, PackageName), TypeNameResourceDiff, nil),
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
			Name: fmt.Sprintf("%s.%s", "incorrect/path", TypeNameResourceDiff),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage("incorrect/path", "path"), TypeNameResourceDiff, nil),
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
			got := IsTypeResourceDiff(testCase.InputType)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
