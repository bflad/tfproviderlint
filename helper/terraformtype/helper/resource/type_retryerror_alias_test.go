package resource

import (
	"go/token"
	"go/types"
	"testing"
)

func TestIsTypeRetryErrorAliases(t *testing.T) {
	pkg := types.NewPackage("github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource", "resource")
	named := types.NewNamed(types.NewTypeName(token.NoPos, pkg, "RetryError", nil), types.NewStruct(nil, nil), nil)
	alias := types.NewAlias(types.NewTypeName(token.NoPos, nil, "RetryAlias", nil), named)
	other := types.NewNamed(types.NewTypeName(token.NoPos, types.NewPackage("example.com/other", "other"), "RetryError", nil), types.NewStruct(nil, nil), nil)
	for _, tc := range []struct {
		name string
		typ  types.Type
		want bool
	}{
		{"named", named, true},
		{"alias", alias, true},
		{"pointer to alias", types.NewPointer(alias), true},
		{"alias to pointer", types.NewAlias(types.NewTypeName(token.NoPos, nil, "PointerAlias", nil), types.NewPointer(named)), true},
		{"unrelated alias", types.NewAlias(types.NewTypeName(token.NoPos, nil, "OtherAlias", nil), other), false},
		{"basic type", types.Typ[types.String], false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsTypeRetryError(tc.typ); got != tc.want {
				t.Errorf("IsTypeRetryError() = %v, want %v", got, tc.want)
			}
		})
	}
}
