package terraformtype

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const (
	FuncParallelTest = `ParallelTest`
	FuncTest         = `Test`

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

// HelperResourceTestCaseInfo represents all gathered TestCase data for easier access
type HelperResourceTestCaseInfo struct {
	AstCompositeLit *ast.CompositeLit
	Fields          map[string]*ast.KeyValueExpr
	TestCase        *resource.TestCase
	TypesInfo       *types.Info
}

// NewHelperResourceTestCaseInfo instantiates a HelperResourceTestCaseInfo
func NewHelperResourceTestCaseInfo(cl *ast.CompositeLit, info *types.Info) *HelperResourceTestCaseInfo {
	result := &HelperResourceTestCaseInfo{
		AstCompositeLit: cl,
		Fields:          astCompositeLitFields(cl),
		TestCase:        &resource.TestCase{},
		TypesInfo:       info,
	}

	return result
}

// DeclaresField returns true if the field name is present in the AST
func (info *HelperResourceTestCaseInfo) DeclaresField(fieldName string) bool {
	return info.Fields[fieldName] != nil
}

// HelperResourceTestStepInfo represents all gathered TestStep data for easier access
type HelperResourceTestStepInfo struct {
	AstCompositeLit *ast.CompositeLit
	Fields          map[string]*ast.KeyValueExpr
	TestStep        *resource.TestStep
	TypesInfo       *types.Info
}

// NewHelperResourceTestStepInfo instantiates a HelperResourceTestStepInfo
func NewHelperResourceTestStepInfo(cl *ast.CompositeLit, info *types.Info) *HelperResourceTestStepInfo {
	result := &HelperResourceTestStepInfo{
		AstCompositeLit: cl,
		Fields:          astCompositeLitFields(cl),
		TestStep:        &resource.TestStep{},
		TypesInfo:       info,
	}

	return result
}

// DeclaresField returns true if the field name is present in the AST
func (info *HelperResourceTestStepInfo) DeclaresField(fieldName string) bool {
	return info.Fields[fieldName] != nil
}

// IsHelperResourceFunc returns if the function call is in the resource package
func IsHelperResourceFunc(e ast.Expr, info *types.Info, funcName string) bool {
	return isPackageFunc(e, info, TypePackagePathHelperResource, funcName)
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
