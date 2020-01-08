package terraformtype

import (
	"go/types"
	"strings"
)

const (
	PackagePathHelperSchema = `github.com/hashicorp/terraform-plugin-sdk/helper/schema`
)

// IsHelperSchemaNamedType returns if the type name matches and is from the helper/schema package
func IsHelperSchemaNamedType(t *types.Named, typeName string) bool {
	if t.Obj().Name() != typeName {
		return false
	}

	// HasSuffix here due to vendoring
	return strings.HasSuffix(t.Obj().Pkg().Path(), PackagePathHelperSchema)
}
