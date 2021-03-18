# tfproviderlint

Static analysis libraries and tooling for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bflad/tfproviderlint)](https://pkg.go.dev/github.com/bflad/tfproviderlint)

## Install

### Local Install

Release binaries are available in the [Releases](https://github.com/bflad/tfproviderlint/releases) section.

To instead use Go to install into your `$GOBIN` directory (e.g. `$GOPATH/bin`):

```console
$ go get github.com/bflad/tfproviderlint/cmd/tfproviderlint
```

If you wish to install the command which includes all linting checks, including [Extra Lint Checks](#extra-lint-checks):

```console
$ go get github.com/bflad/tfproviderlint/cmd/tfproviderlintx
```

### Docker Install

```console
$ docker pull bflad/tfproviderlint
```

### Homebrew Install

```console
$ brew install bflad/tap/tfproviderlint
```

## Usage

The `tfproviderlint` and `tfproviderlintx` tools operate similarly except for which checks are available. Additional information about usage and configuration options can be found by passing the `help` argument:

```console
$ tfproviderlint help
```

To enable only specific checks, they can be passed in as flags:

```console
$ tfproviderlint -AT001
```

To enable all checks, but disable specific checks, they can be passed in as flags set to `false`:

```console
$ tfproviderlint -AT001=false
```

### Local Usage

To report issues, change into the directory of the Terraform Provider code and run:

```console
$ tfproviderlint ./...
```

To apply automated fixes for checks that support them, change into the directory of the Terraform Provider code and run:

```console
$ tfproviderlint -fix ./...
```

It is also possible to run via [`go vet`](https://golang.org/cmd/vet/):

```console
$ go vet -vettool $(which tfproviderlint) ./...
```

### Docker Usage

Change into the directory of the Terraform Provider code and run:

```console
$ docker run -v $(pwd):/src bflad/tfproviderlint ./...
```

### GitHub Action Usage

A [GitHub Action](https://github.com/features/actions) is available: [tfproviderlint-github-action](https://github.com/bflad/tfproviderlint-github-action)

## Standard Lint Checks

Standard lint checks are enabled by default in the `tfproviderlint` tool. Opt-in checks can be found in the [Extra Lint Checks section](#extra-lint-checks). For additional information about each check, you can run `tfproviderlint help NAME`.

### Standard Acceptance Test Checks

| Check | Description | Type |
|---|---|---|
| [AT001](passes/AT001) | check for `TestCase` missing `CheckDestroy` | AST |
| [AT002](passes/AT002) | check for acceptance test function names including the word import | AST |
| [AT003](passes/AT003) | check for acceptance test function names missing an underscore | AST |
| [AT004](passes/AT004) | check for `TestStep` `Config` containing provider configuration | AST |
| [AT005](passes/AT005) | check for acceptance test function names missing `TestAcc` prefix | AST |
| [AT006](passes/AT006) | check for acceptance test functions containing multiple `resource.Test()` invocations | AST |
| [AT007](passes/AT007) | check for acceptance test functions containing multiple `resource.ParallelTest()` invocations | AST |
| [AT008](passes/AT008) | check for acceptance test function declaration `*testing.T` parameter naming | AST |
| [AT009](passes/AT009) | check for `acctest.RandStringFromCharSet()` calls that can be simplified to `acctest.RandString()` | AST |

### Standard Resource Checks

| Check | Description | Type |
|---|---|---|
| [R001](passes/R001) | check for `ResourceData.Set()` calls using complex key argument | AST |
| [R002](passes/R002) | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| [R003](passes/R003) | check for `Resource` having `Exists` functions | AST |
| [R004](passes/R004) | check for `ResourceData.Set()` calls using incompatible value types | AST |
| [R005](passes/R005) | check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call | AST |
| [R006](passes/R006) | check for `RetryFunc` that omit retryable errors | AST |
| [R007](passes/R007) | check for deprecated `(schema.ResourceData).Partial` usage | AST |
| [R008](passes/R008) | check for deprecated `(schema.ResourceData).SetPartial` usage | AST |
| [R009](passes/R009) | check for Go panic usage | AST |
| [R010](passes/R010) | check for `(schema.ResourceData).GetChange` assignment which should use `(schema.ResourceData).Get` | AST |
| [R011](passes/R011) | check for `Resource` that configure `MigrateState` | AST |
| [R012](passes/R012) | check for data source `Resource` that configure `CustomizeDiff` | AST |
| [R013](passes/R013) | check for `map[string]*Resource` that resource names contain at least one underscore | AST |
| [R014](passes/R014) | check for `CreateFunc`, `CreateContextFunc`, `DeleteFunc`, `DeleteContextFunc`, `ReadFunc`, `ReadContextFunc`, `UpdateFunc`, and `UpdateContextFunc` parameter naming | AST |
| [R015](passes/R015) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.UniqueId()` value | AST |
| [R016](passes/R016) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `resource.PrefixedUniqueId()` value | AST |
| [R017](passes/R017) | check for `(*schema.ResourceData).SetId()` receiver method usage with unstable `time.Now()` value | AST |
| [R018](passes/R018) | check for `time.Sleep()` function usage | AST |
| [R019](passes/R019) | check for `(*schema.ResourceData).HasChanges()` receiver method usage with many arguments | AST |

### Standard Schema Checks

| Check | Description | Type |
|---|---|---|
| [S001](passes/S001) | check for `Schema` of `TypeList` or `TypeSet` missing `Elem` | AST |
| [S002](passes/S002) | check for `Schema` with both `Required` and `Optional` enabled | AST |
| [S003](passes/S003) | check for `Schema` with both `Required` and `Computed` enabled | AST |
| [S004](passes/S004) | check for `Schema` with both `Required` and `Default` configured | AST |
| [S005](passes/S005) | check for `Schema` with both `Computed` and `Default` configured | AST |
| [S006](passes/S006) | check for `Schema` of `TypeMap` missing `Elem` | AST |
| [S007](passes/S007) | check for `Schema` with both `Required` and `ConflictsWith` configured | AST |
| [S008](passes/S008) | check for `Schema` of `TypeList` or `TypeSet` with `Default` configured | AST |
| [S009](passes/S009) | check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` configured | AST |
| [S010](passes/S010) | check for `Schema` of `Computed` only with `ValidateFunc` configured | AST |
| [S011](passes/S011) | check for `Schema` of `Computed` only with `DiffSuppressFunc` configured | AST |
| [S012](passes/S012) | check for `Schema` that `Type` is configured | AST |
| [S013](passes/S013) | check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured | AST |
| [S014](passes/S014) | check for `Schema` within `Elem` that `Computed`, `Optional`, and `Required` are not configured | AST |
| [S015](passes/S015) | check for `map[string]*Schema` that attribute names are valid | AST |
| [S016](passes/S016) | check for `Schema` that `Set` is only configured for `TypeSet` | AST |
| [S017](passes/S017) | check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet` | AST |
| [S018](passes/S018) | check for `Schema` that should use `TypeList` with `MaxItems: 1` | AST |
| [S019](passes/S019) | check for `Schema` that should omit `Computed`, `Optional`, or `Required` set to `false` | AST |
| [S020](passes/S020) | check for `Schema` of `Computed` only with `ForceNew` enabled | AST |
| [S021](passes/S021) | check for `Schema` that should omit `ComputedWhen` | AST |
| [S022](passes/S022) | check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource` | AST |
| [S023](passes/S023) | check for `Schema` that should omit `Elem` with incompatible `Type` | AST |
| [S024](passes/S024) | check for `Schema` that should omit `ForceNew` in data source schema attributes | AST |
| [S025](passes/S025) | check for `Schema` of `Computed` only with `AtLeastOneOf` configured | AST |
| [S026](passes/S026) | check for `Schema` of `Computed` only with `ConflictsWith` configured | AST |
| [S027](passes/S027) | check for `Schema` of `Computed` only with `Default` configured | AST |
| [S028](passes/S028) | check for `Schema` of `Computed` only with `DefaultFunc` configured | AST |
| [S029](passes/S029) | check for `Schema` of `Computed` only with `ExactlyOneOf` configured | AST |
| [S030](passes/S030) | check for `Schema` of `Computed` only with `InputDefault` configured | AST |
| [S031](passes/S031) | check for `Schema` of `Computed` only with `MaxItems` configured | AST |
| [S032](passes/S032) | check for `Schema` of `Computed` only with `MinItems` configured | AST |
| [S033](passes/S033) | check for `Schema` of `Computed` only with `StateFunc` configured | AST |
| [S034](passes/S034) | check for `Schema` that configure `PromoteSingle` | AST |
| [S035](passes/S035) | check for `Schema` with invalid `AtLeastOneOf` attribute references | AST |
| [S036](passes/S036) | check for `Schema` with invalid `ConflictsWith` attribute references | AST |
| [S037](passes/S037) | check for `Schema` with invalid `ExactlyOneOf` attribute references | AST |

### Standard Validation Checks

| Check | Description | Type |
|---|---|---|
| [V001](passes/V001) | check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()` | AST |
| [V002](passes/V002) | check for deprecated `CIDRNetwork` validation function usage | AST |
| [V003](passes/V003) | check for deprecated `IPRange` validation function usage | AST |
| [V004](passes/V004) | check for deprecated `SingleIP` validation function usage | AST |
| [V005](passes/V005) | check for deprecated `ValidateJsonString` validation function usage | AST |
| [V006](passes/V006) | check for deprecated `ValidateListUniqueStrings` validation function usage | AST |
| [V007](passes/V007) | check for deprecated `ValidateRegexp` validation function usage | AST |
| [V008](passes/V008) | check for deprecated `ValidateRFC3339TimeString` validation function usage | AST |
| [V009](passes/V009) | check for `validation.StringMatch()` call with empty message argument | AST |
| [V010](passes/V010) | check for `validation.StringDoesNotMatch()` call with empty message argument | AST |

## Extra Lint Checks

Extra lint checks are not included in the `tfproviderlint` tool and must be accessed via the `tfproviderlintx` tool or [added to a custom lint tool](#implementing-a-custom-lint-tool). Generally these represent advanced Terraform Plugin SDK functionality that is not appropriate for all Terraform Providers.

### Extra Acceptance Test Checks

| Check | Description | Type |
|---|---|---|
| [XAT001](xpasses/XAT001) | check for `TestCase` missing `ErrorCheck` | AST |

### Extra Resource Checks

| Check | Description | Type |
|---|---|---|
| [XR001](xpasses/XR001) | check for usage of `ResourceData.GetOkExists()` calls | AST |
| [XR002](xpasses/XR002) | check for `Resource` that should implement `Importer` | AST |
| [XR003](xpasses/XR003) | check for `Resource` that should implement `Timeouts` | AST |
| [XR004](xpasses/XR004) | check for `ResourceData.Set()` calls that should implement error checking with complex values | AST |
| [XR005](xpasses/XR005) | check for `Resource` that `Description` is configured | AST |
| [XR006](xpasses/XR006) | check for `Resource` that implements `Timeouts` for missing `Create`, `Delete`, `Read`, or `Update` implementation | AST |

### Extra Schema Checks

| Check | Description | Type |
|---|---|---|
| [XS001](xpasses/XS001) | check for `map[string]*Schema` that `Description` is configured | AST |
| [XS002](xpasses/XS002) | check for `map[string]*Schema` that keys are in alphabetical order | AST |

## Development and Testing

This project is built on the [`go/analysis`](https://godoc.org/golang.org/x/tools/go/analysis) framework and uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

Helpful tooling for development:

* [`astdump`](https://github.com/wingyplus/astdump): a tool for displaying the AST form of Go file
* [`ssadump`](https://godoc.org/golang.org/x/tools/cmd/ssadump): a tool for displaying and interpreting the SSA form of Go programs

### Adding an Analyzer

* Create new analyzer in `passes/` (or `xpasses/` for extra checks)
* If the `Analyzer` reports issues, add to `AllChecks` variable in `passes/checks.go` (or `xpasses/checks.go` for extra checks)
* Since the [`analysistest` package](https://godoc.org/golang.org/x/tools/go/analysis/analysistest) does not support Go Modules currently, each analyzer that implements testing must add a symlink to the top level `vendor` directory in the `testdata/src/a` directory. e.g. `ln -s ../../../../../vendor passes/NAME/testdata/src/a/vendor`

### Implementing SuggestedFixes Testing

The upstream `analysistest` package now contains functionality to verify `SuggestedFixes` via `RunWithSuggestedFixes`.

```go
import (
  "testing"

  "golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerFixes(t *testing.T) {
  testdata := analysistest.TestData()
  analysistest.RunWithSuggestedFixes(t, testdata, Analyzer, "a")
}
```

To setup the expected file content verification, the testing expects a file suffixed with `.golden` (e.g. `testdata/src/a/main.go.golden`).

### Implementing a Custom Lint Tool

The `go/analysis` framework and this codebase are designed for flexibility. You may wish to permanently disable certain default checks or even implement your own provider-specific checks. An example of how to incorporate all default and extra checks in a CLI command can be found in `cmd/tfproviderlintx`. To permanently exclude checks, each desired `Analyzer` must be individually included, similar to how `AllChecks()` is built in `passes/checks.go`.

The `passes` directory also includes the underlying `Analyzer` which iteratively gather AST-based information about the Terraform Provider code being analyzed. For example, `passes/helper/resource/retryfuncinfo` returns information from all named and anonymous declarations of `helper/resource.RetryFunc()`.

Primatives for working with Terraform Plugin SDK primatives can be found in `helper/terraformtype`. Primatives for working with the Go AST can be found in `helper/astutils`.

### Updating Dependencies

```console
$ go get URL
$ go mod tidy
$ go mod vendor
```

### Unit Testing

```console
$ go test ./...
```

### Local Install Testing

```console
$ go install ./cmd/tfproviderlint
```
