package terraformtype

import (
	"go/types"
)

const (
	PackagePathHelperSchema = `github.com/hashicorp/terraform-plugin-sdk/helper/schema`
)

// IsHelperSchemaNamedType returns if the type name matches and is from the helper/schema package
func IsHelperSchemaNamedType(t *types.Named, typeName string) bool {
	return isPackageNamedType(t, PackagePathHelperSchema, typeName)
}
