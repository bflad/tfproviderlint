# v0.7.0

FEATURES

* **New Analyzer:** `schemavalidatefunc`: returns `SchemaValidateFunc` declarations
* **New Check:** `R005`: check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call
* **New Check:** `S021`: check for `Schema` that should omit `ComputedWhen`
* **New Check:** `S022`: check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource`
* **New Check:** `S023`: check for `Schema` that should omit `Elem` with incompatible `Type`
* **New Check:** `V001`: check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()`

ENHANCEMENTS

* cmd/tfproviderlint: Add `-V` and `-version` flags for version information (#40)
* helper/astutils: Functions for determining package functions, package receiver methods, package types, and some function parameter types
* helper/terraformtype: Support pointers with `IsHelperResourceTypeTestCase()`, `IsHelperResourceTypeTestStep()`, `IsHelperSchemaTypeResource()`, `IsHelperSchemaTypeResourceData()`, `IsHelperSchemaTypeSchema()`, and `IsHelperSchemaTypeSet()` functions
* passes: Add `AllChecks` variable which can be used to bootstrap custom downstream linters

# v0.6.0

BREAKING CHANGES

* The `acctestcase`, `schemaresource`, and `schemaschema` passes now return results with `helper/terraformtype` types (`[]*terraformtype.HelperResourceTestCaseInfo`, `[]*terraformtype.HelperSchemaResourceInfo`, and `[]*terraformtype.HelperSchemaSchemaInfo` respectively) instead of `[]*ast.CompositeLit`. The `*ast.CompositeLit` can be accessed via the `AstCompositeLit` field of each of the new types.

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
