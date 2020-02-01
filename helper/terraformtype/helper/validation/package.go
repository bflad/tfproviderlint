package validation

import (
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	PackageName = `validation`
	PackagePath = `github.com/hashicorp/terraform-plugin-sdk/helper/validation`
)

// IsPackageFunc returns if the function call is in the package
func IsPackageFunc(e ast.Expr, info *types.Info, funcName string) bool {
	return astutils.IsPackageFunc(e, info, PackagePath, funcName)
}

// IsPackageNamedType returns if the type name matches and is from the package
func IsPackageNamedType(t *types.Named, typeName string) bool {
	return astutils.IsPackageNamedType(t, PackagePath, typeName)
}
