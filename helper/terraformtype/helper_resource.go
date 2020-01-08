package terraformtype

import (
	"go/ast"
	"go/types"
)

const (
	FuncNameParallelTest = `ParallelTest`
	FuncNameTest         = `Test`

	PackagePathHelperResource = `github.com/hashicorp/terraform-plugin-sdk/helper/resource`
)

// IsHelperResourceFunc returns if the function call is in the resource package
func IsHelperResourceFunc(e ast.Expr, info *types.Info, funcName string) bool {
	return isPackageFunc(e, info, PackagePathHelperResource, funcName)
}

// IsHelperResourceNamedType returns if the type name matches and is from the helper/resource package
func IsHelperResourceNamedType(t *types.Named, typeName string) bool {
	return isPackageNamedType(t, PackagePathHelperResource, typeName)
}
