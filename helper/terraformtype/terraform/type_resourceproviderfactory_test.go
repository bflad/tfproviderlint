package terraform

import (
	"fmt"
	"go/token"
	"go/types"
	"testing"
)

func TestIsTypeResourceProviderFactory(t *testing.T) {
	testCases := []struct {
		Name      string
		InputType types.Type
		Expected  bool
	}{
		{
			Name: fmt.Sprintf("%s.%s", PackagePathVersion(1), TypeNameResourceProviderFactory),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePathVersion(1), PackageName), TypeNameResourceProviderFactory, nil),
				nil,
				nil,
			),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("*%s.%s", PackagePathVersion(1), TypeNameResourceProviderFactory),
			InputType: types.NewPointer(types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePathVersion(1), PackageName), TypeNameResourceProviderFactory, nil),
				nil,
				nil,
			)),
			Expected: true,
		},
		{
			Name: fmt.Sprintf("%s.%s", PackagePathVersion(1), "Not"),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage(PackagePathVersion(1), PackageName), "Not", nil),
				nil,
				nil,
			),
			Expected: false,
		},
		{
			Name: fmt.Sprintf("%s.%s", "incorrect/path", TypeNameResourceProviderFactory),
			InputType: types.NewNamed(
				types.NewTypeName(token.NoPos, types.NewPackage("incorrect/path", "path"), TypeNameResourceProviderFactory, nil),
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
			got := IsTypeResourceProviderFactory(testCase.InputType)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
