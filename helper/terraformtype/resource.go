package terraformtype

import (
	"go/ast"
	"go/types"
	"strings"
)

const (
	TestCaseFieldCheckDestroy              = `CheckDestroy`
	TestCaseFieldIDRefreshName             = `IDRefreshName`
	TestCaseFieldIDRefreshIgnore           = `IDRefreshIgnore`
	TestCaseFieldIsUnitTest                = `IsUnitTest`
	TestCaseFieldPreCheck                  = `PreCheck`
	TestCaseFieldPreventPostDestroyRefresh = `PreventPostDestroyRefresh`
	TestCaseFieldProviders                 = `Providers`
	TestCaseFieldProviderFactories         = `ProviderFactories`
	TestCaseFieldSteps                     = `Steps`

	TestStepFieldCheck                     = `Check`
	TestStepFieldConfig                    = `Config`
	TestStepFieldDestroy                   = `Destroy`
	TestStepFieldExpectError               = `ExpectError`
	TestStepFieldExpectNonEmptyPlan        = `ExpectNonEmptyPlan`
	TestStepFieldImportState               = `ImportState`
	TestStepFieldImportStateId             = `ImportStateId`
	TestStepFieldImportStateIdFunc         = `ImportStateIdFunc`
	TestStepFieldImportStateIdPrefix       = `ImportStateIdPrefix`
	TestStepFieldImportStateCheck          = `ImportStateCheck`
	TestStepFieldImportStateVerify         = `ImportStateVerify`
	TestStepFieldImportStateVerifyIgnore   = `ImportStateVerifyIgnore`
	TestStepFieldPlanOnly                  = `PlanOnly`
	TestStepFieldPreConfig                 = `PreConfig`
	TestStepFieldPreventDiskCleanup        = `PreventDiskCleanup`
	TestStepFieldPreventPostDestroyRefresh = `PreventPostDestroyRefresh`
	TestStepFieldResourceName              = `ResourceName`
	TestStepFieldSkipFunc                  = `SkipFunc`
	TestStepFieldTaint                     = `Taint`

	TypeNameTestCase = `TestCase`
	TypeNameTestStep = `TestStep`

	TypePackagePathHelperResource = `github.com/hashicorp/terraform-plugin-sdk/helper/resource`
)

// HelperResourceTypeTestCaseContainsFields checks if any fields are present in the TestCase
func HelperResourceTypeTestCaseContainsFields(testCase *ast.CompositeLit, fields ...string) bool {
	return astCompositeLitContainsAnyField(testCase, fields...)
}

// HelperResourceTypeTestCaseCheckDestroy extracts the TestCase CheckDestroy value
// NOTE: The current function signature output type should not be considered stable
//       except for being a pointer.
func HelperResourceTypeTestCaseCheckDestroy(testCase *ast.CompositeLit) *ast.Expr {
	return astCompositeLitFieldExprValue(testCase, TestCaseFieldCheckDestroy)
}

// IsHelperResourceTypeTestCase returns if the type is TestCase from the helper/schema package
func IsHelperResourceTypeTestCase(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsHelperResourceNamedType(t, TypeNameTestCase)
	default:
		return false
	}
}

// IsHelperResourceNamedType returns if the type name matches and is from the helper/resource package
func IsHelperResourceNamedType(t *types.Named, typeName string) bool {
	if t.Obj().Name() != typeName {
		return false
	}

	// HasSuffix here due to vendoring
	if !strings.HasSuffix(t.Obj().Pkg().Path(), TypePackagePathHelperResource) {
		return false
	}

	return true
}
