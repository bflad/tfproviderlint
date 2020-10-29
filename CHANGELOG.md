# v0.20.0

FEATURES

* **New Check:** `AT009`: check for `acctest.RandStringFromCharSet()` calls that can be `acctest.RandString`

ENHANCEMENTS

* passes/helper/acctest: Pass for collecting `acctest.RandStringFromCharSet()` calls
* passes/АТ001: Add optional `-ignored-filename-suffixes` and `-ignored-filename-prefixes` flags (#210)

# v0.19.0

FEATURES

* **New Check:** `R018`: check for `time.Sleep()` function usage

ENHANCEMENTS

* passes/commentignore: Support multiple comma separated keys and extra `//` comments (#208)
* passes/stdlib: Pass for collecting `time.Sleep()` calls

BUG FIXES

* helper/astutils: Prevent `(*types.Info).ObjectOf()` type assertion panics (#207)

# v0.18.0

ENHANCEMENTS

* passes/helper/validation: Pass for collecting `validation.StringInSlice()` calls
* passes/stdlib: Pass for collecting `fmt.Errorf()` calls

# v0.17.0

FEATURES

* **New Check:** `R015`: check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.UniqueId()` value
* **New Check:** `R016`: check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.PrefixedUniqueId()` value
* **New Check:** `R017`: check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `time.Now()` value

ENHANCEMENTS

* helper/analysisutils: Add `StdlibFunctionCallExprAnalyzer` with associated runner
* helper/astutils: Support standard library handling equivalents (no vendoring) for package functions
* passes/helper/schema: Pass for collecting `(*schema.ResourceData).SetId()` calls
* passes/stdlib: Pass for collecting `fmt.Sprintf()` calls

# v0.16.0

BREAKING CHANGES

* helper/terraformtype/helper/resource: Real SDK v1 types inside `RetryErrorInfo`, `TestCaseInfo`, and `TestStepInfo` have been replaced with internal representations to prevent importing SDK versions.
* helper/terraformtype/helper/schema: Real SDK v1 types inside `SchemaInfo` and `ResourceInfo` have been replaced with internal representations to prevent importing SDK versions.

FEATURES

* **New Extra Check:** `XS002`: check for `map[string]*Schema` that keys are in alphabetical order (#177)

# v0.15.0

BREAKING CHANGES

* helper/analysisfixtest: Removed package since testing for `SuggestedFixes` is now available with `golang.org/x/tools/go/analysis/analysistest.RunWithSuggestedFixes`

FEATURES

* **New Extra Check:** `XR005`: check for `Resource` that `Description` is configured

ENHANCEMENTS

* helper/terraformtype/helper/schema: Support basic linking of `ResourceInfo` and `SchemaInfo` (#129)
* passes: Initial support for Terraform Plugin SDK v2
* passes/terraform: Pass for collecting `terraform.ResourceProviderFactory` selectors
* xpasses: Initial support for Terraform Plugin SDK v2

# v0.14.0

ENHANCEMENTS

* helper/analysisutils: Add `DeprecatedWithReplacementPointerSelectorExprAnalyzer`, `TypeAssertExprRemovalAnalyzer` and `TypeAssertExprAnalyzer` with associated runners
* helper/astutils: Support `*ast.StarExpr` in `IsPackageFunc`
* helper/terraformtype/helper/schema: Constants and functions for working with `Provider` type
* helper/terraformtype/terraform: Constants and functions for working with `ResourceProviderFactory` type
* passes/helper/schema: Pass for collecting `*schema.Provider` type assertions
* passes/terraform: Pass for collecting `terraform.ResourceProvider` selectors

# v0.13.0

BREAKING CHANGES

* helper/analysisutils: `DeprecatedReceiverMethodSelectorExprAnalyzer` and `DeprecatedReceiverMethodSelectorExprRunner` now expect an `*ast.CallExpr` analyzer and package path (e.g. `github.com/hashicorp/terraform-plugin-sdk/helper/validation`) instead of just package name (e.g. `validation`)
* helper/analysisutils: `DeprecatedWithReplacementSelectorExprAnalyzer` and `DeprecatedWithReplacementSelectorExprRunner` now expect package paths (e.g. `github.com/hashicorp/terraform-plugin-sdk/helper/validation`) instead of just package names (e.g. `validation`)

ENHANCEMENTS

* helper/analysisutils: `DeprecatedReceiverMethodSelectorExprRunner` now implements `SuggestedFixes`
* helper/analysisutils: `DeprecatedWithReplacementSelectorExprRunner` now implements `SuggestedFixes`
* helper/astutils: Add `HasFieldListLength`, `IsExprTypeBool`, `IsExprTypeMapStringInterface` functions
* helper/terraformtype/helper/customdiff: Constants and functions for working with entire package
* helper/terraformtype/helper/schema: Constants and functions for working with `CustomizeDiffFunc` type
* helper/terraformtype/helper/schema: Constants and functions for working with `StateUpgradeFunc` type
* helper/terraformtype/terraform: Constants and functions for working with `ResourceProvider` type
* passes/helper/customdiff: Passes for collecting `customdiff` function information for further analysis
* passes/helper/schema: Pass for collecting `CustomizeDiffFunc` function information for further analysis
* passes/helper/schema: Pass for collecting `StateUpgradeFunc` function information for further analysis
* passes/R007: Support suggested fix
* passes/R008: Support suggested fix
* passes/V002: Support suggested fix
* passes/V003: Support suggested fix
* passes/V004: Support suggested fix
* passes/V005: Support suggested fix
* passes/V006: Support suggested fix
* passes/V007: Support suggested fix
* passes/V008: Support suggested fix

# v0.12.0

BREAKING CHANGES

* helper/astutils: The `IsFunctionParameterType*` functions have been renamed `IsExprType*` to better denote their purpose.

FEATURES

* **New Check:** `S035`: check for `Schema` with invalid `AtLeastOneOf` attribute references
* **New Check:** `S036`: check for `Schema` with invalid `ConflictsWith` attribute references
* **New Check:** `S037`: check for `Schema` with invalid `ExactlyOneOf` attribute references

# v0.11.0

ENHANCEMENTS

* passes/AT004: Support ignore comments
* passes/AT005: Support ignore comments
* passes/AT006: Support ignore comments
* passes/AT007: Support ignore comments

# v0.10.0

BREAKING CHANGES

* helper/terraformtype/helper/schema: `NewSchemaValidateFuncInfo` now accepts a single `ast.Node` instead of separate `*ast.FuncDecl` and `*ast.FuncLit`

FEATURES

* **New Check:** `AT008`: check for acceptance test function declaration `*testing.T` parameter naming
* **New Check:** `R007`: check for deprecated `(schema.ResourceData).Partial` receiver method usage
* **New Check:** `R008`: check for deprecated `(schema.ResourceData).SetPartial` receiver method usage
* **New Check:** `R009`: check for Go panic usage
* **New Check:** `R010`: check for `(schema.ResourceData).GetChange` assignment which should use `(schema.ResourceData).Get`
* **New Check:** `R011`: check for `Resource` that configure `MigrateState`
* **New Check:** `R012`: check for data source `Resource` that configure `CustomizeDiff`
* **New Check:** `R013`: check for `map[string]*Resource` that resource names contain at least one underscore
* **New Check:** `R014`: check for `CreateFunc`, `DeleteFunc`, `ReadFunc`, and `UpdateFunc` parameter naming
* **New Check:** `S034`: check for `Schema` that configure `PromoteSingle`

# v0.9.0

BREAKING CHANGES

* helper/terraformtype: Refactored into helper/resource and helper/schema subpackages
* passes/schemamap: `GetSchemaAttributes` has been replaced with `helper/terraformtype/helper/schema.GetSchemaMapSchemas`
* passes/schemamap: `GetSchemaAttributeNames` has been replaced with `helper/terraformtype/helper/schema.GetSchemaMapAttributeNames`

FEATURES

* **New Check:** `S024`: check for `Schema` that should omit `ForceNew` in data source schema attributes
* **New Check:** `S025`: check for `Schema` of only `Computed` enabled with `AtLeastOneOf` configured
* **New Check:** `S026`: check for `Schema` of only `Computed` enabled with `ConflictsWith` configured
* **New Check:** `S027`: check for `Schema` of only `Computed` enabled with `Default` configured
* **New Check:** `S028`: check for `Schema` of only `Computed` enabled with `DefaultFunc` configured
* **New Check:** `S029`: check for `Schema` of only `Computed` enabled with `ExactlyOneOf` configured
* **New Check:** `S030`: check for `Schema` of only `Computed` enabled with `InputDefault` configured
* **New Check:** `S031`: check for `Schema` of only `Computed` enabled with `MaxItems` configured
* **New Check:** `S032`: check for `Schema` of only `Computed` enabled with `MinItems` configured
* **New Check:** `S033`: check for `Schema` of only `Computed` enabled with `StateFunc` configured
* **New Check:** `V002`: check for deprecated `CIDRNetwork` validation function usage
* **New Check:** `V003`: check for deprecated `IPRange` validation function usage
* **New Check:** `V004`: check for deprecated `SingleIP` validation function usage
* **New Check:** `V005`: check for deprecated `ValidateJsonString` validation function usage
* **New Check:** `V006`: check for deprecated `ValidateListUniqueStrings` validation function usage
* **New Check:** `V007`: check for deprecated `ValidateRegexp` validation function usage
* **New Check:** `V008`: check for deprecated `ValidateRFC3339TimeString` validation function usage

ENHANCEMENTS:

* helper/analysisutils: `DeprecatedWithReplacementSelectorExprAnalyzer`, `FunctionCallExprAnalyzer`, `ReceiverMethodCallExprAnalyzer`, and `SelectorExprAnalyzer` Analyzer generators
* helper/astutils: `ExprBoolValue`, `ExprIntValue`, and `ExprStringValue` functions
* helper/terraformtype/helper/schema: `GetSchemaMapAttributeNames`, `GetSchemaMapSchemas`, and `IsMapStringHelperSchemaTypeSchema` functions
* helper/terraformtype/helper/schema: `ResourceInfo` type `IsDataSource` and `IsResource` receiver methods

# v0.8.0

FEATURES

* **New Analyzer:** `resourcedataget`: returns `ResourceData.Get()` calls for later passes
* **New Analyzer:** `resourcedatagetchange`: returns `ResourceData.GetChange()` calls for later passes
* **New Analyzer:** `resourcedatagetok`: returns `ResourceData.GetOk()` calls for later passes
* **New Analyzer:** `resourcedatagetokexists`: returns `ResourceData.GetOkExists()` calls for later passes
* **New Command:** `tfproviderlintx`: Runs all standard and extra checks
* **New Extra Check:** `XS001`: check for `map[string]*Schema` that `Description` is configured
* **New Extra Check:** `XR001`: check for usage of `ResourceData.GetOkExists()` calls
* **New Extra Check:** `XR002`: check for `Resource` that should implement `Importer`
* **New Extra Check:** `XR003`: check for `Resource` that should implement `Timeouts`
* **New Extra Check:** `XR004`: check for `ResourceData.Set()` calls that should implement error checking with complex values

# v0.7.0

FEATURES

* **New Analyzer:** `retryfunc`: returns `RetryFunc` declarations
* **New Analyzer:** `schemavalidatefunc`: returns `SchemaValidateFunc` declarations
* **New Check:** `R005`: check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call
* **New Check:** `R006`: check for `RetryFunc` that omit retryable errors
* **New Check:** `S021`: check for `Schema` that should omit `ComputedWhen`
* **New Check:** `S022`: check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource`
* **New Check:** `S023`: check for `Schema` that should omit `Elem` with incompatible `Type`
* **New Check:** `V001`: check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()`

ENHANCEMENTS

* cmd/tfproviderlint: Add `-V` and `-version` flags for version information (#40)
* helper/astutils: Functions for determining package functions, package receiver methods, package types, and some function parameter types
* helper/terraformtype: Support pointers with `IsHelperResourceTypeTestCase()`, `IsHelperResourceTypeTestStep()`, `IsHelperSchemaTypeResource()`, `IsHelperSchemaTypeResourceData()`, `IsHelperSchemaTypeSchema()`, and `IsHelperSchemaTypeSet()` functions
* helper/terraformtype: Add constants for helper/resource `NonRetryableError` and `RetryableError` function names
* helper/terraformtype: Add constants and helper functions for helper/resource `RetryError` type
* passes: Add `AllChecks` variable which can be used to bootstrap custom downstream linters

# v0.6.0

BREAKING CHANGES

* The `acctestcase`, `schemaresource`, and `schemaschema` passes now return results with `helper/terraformtype` types (`[]*terraformtype.TestCaseInfo`, `[]*terraformtype.HelperSchemaResourceInfo`, and `[]*terraformtype.HelperSchemaSchemaInfo` respectively) instead of `[]*ast.CompositeLit`. The `*ast.CompositeLit` can be accessed via the `AstCompositeLit` field of each of the new types.

NOTES

* Tests are now Go 1.13 compatible
* Updates github.com/hashicorp/terraform-plugin-sdk dependency to v1.4.1 which cleans up some transitive dependencies

FEATURES

* **New Analyzer:** `accteststep`: returns `resource.TestStep` literals for later passes
* **New Analyzer:** `testcheckresourceattr`: returns `resource.TestCheckResourceAttr()` calls for later passes
* **New Analyzer:** `testcheckresourceattrset`: returns `resource.TestCheckResourceAttrSet()` calls for later passes
* **New Analyzer:** `testmatchresourceattr`: returns `resource.TestMatchResourceAttr()` calls for later passes
* **New Check:** `AT005`: check for acceptance test names missing `TestAcc` prefix
* **New Check:** `AT006`: check for acceptance tests containing multiple `resource.Test()` invocations
* **New Check:** `AT007`: check for acceptance tests containing multiple `resource.ParallelTest()` invocations
* **New Check:** `S020`: check for `Schema` of only `Computed` enabled with `ForceNew` enabled

ENHANCEMENTS

* helper/terraformtype: Create enhanced AST types and field name constants for helper/schema.Resource, helper/schema.Schema, helper/resource.TestCase, and helper/resource.TestStep types

# v0.5.1

NOTES

* Updates github.com/hashicorp/terraform-plugin-sdk dependency to allow repositories using this project as a dependency to switch from github.com/hashicorp/hcl2 to github.com/hashicorp/hcl/v2

# v0.5.0

BREAKING CHANGES

* The import path checking used for the linters and underlying Terraform codebase dependency has been migrated from `github.com/hashicorp/terraform` to `github.com/hashicorp/terraform-plugin-sdk`. For more information see the [Terraform Plugin SDK page in the Extending Terraform documentation](https://www.terraform.io/docs/extend/plugin-sdk.html).

# v0.4.0

FEATURES

* **New Check:** `S007`: check for `Schema` with both `Required` and `ConflictsWith` configured
* **New Check:** `S008`: check for `Schema` of `TypeList` or `TypeSet` with `Default` configured
* **New Check:** `S009`: check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` configured
* **New Check:** `S010`: check for `Schema` of only `Computed` enabled with `ValidateFunc` configured
* **New Check:** `S011`: check for `Schema` of only `Computed` enabled with `DiffSuppressFunc` configured
* **New Check:** `S012`: check for `Schema` that `Type` is configured
* **New Check:** `S013`: check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured
* **New Check:** `S014`: check for `Schema` within `Elem` that `Computed`, `Optional`, or `Required` are not configured
* **New Check:** `S015`: check for `map[string]*Schema` that attribute names are valid
* **New Check:** `S016`: check for `Schema` that `Set` is only configured for `TypeSet`
* **New Check:** `S017`: check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet`
* **New Check:** `S018`: check for `Schema` that should prefer `TypeList` with `MaxItems: 1`
* **New Check:** `S019`: check for `Schema` that should omit `Computed`, `Optional`, or `Required` set to `false`

BUG FIXES

* passes/AT001: Ignore file names beginning with `data_source_` [GH-25]

# v0.3.0

ENHANCEMENTS

* Support `map[string]*schema.Schema` in schema checks [GH-24]

# v0.2.0

FEATURES:

* **New Check** `R004` [GH-21]

# v0.1.0

* Initial release with checks:
  * `AT001`, `AT002`, `AT003`, `AT004`
  * `R001`, `R002`, `R003`
  * `S001`, `S002`, `S003`, `S004`, `S005`, `S006`
