package customdiff

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype"
)

const (
	PackageModule     = terraformtype.ModuleTerraformPluginSdk
	PackageModulePath = `helper/customdiff`
	PackageName       = `customdiff`
	PackagePath       = PackageModule + `/` + PackageModulePath
)

// IsFunc returns if the function call is in the customdiff package
func IsFunc(e ast.Expr, info *types.Info, funcName string) bool {
	return astutils.IsModulePackageFunc(e, info, PackageModule, PackageModulePath, funcName)
}

// IsNamedType returns if the type name matches and is from the helper/customdiff package
func IsNamedType(t *types.Named, typeName string) bool {
	return astutils.IsModulePackageNamedType(t, PackageModule, PackageModulePath, typeName)
}

// PackagePathVersion returns the import path for a module version
func PackagePathVersion(moduleVersion int) string {
	switch moduleVersion {
	case 0, 1:
		return PackagePath
	default:
		return fmt.Sprintf("%s/v%d/%s", PackageModule, moduleVersion, PackageModulePath)
	}
}
