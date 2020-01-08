package terraformtype

import (
	"go/ast"
	"go/types"
	"strings"
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
	if t.Obj().Name() != typeName {
		return false
	}

	// HasSuffix here due to vendoring
	if !strings.HasSuffix(t.Obj().Pkg().Path(), PackagePathHelperResource) {
		return false
	}

	return true
}
