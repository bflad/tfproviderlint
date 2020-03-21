# tfproviderlint

A linting tool for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

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
| [AT001](passes/AT001/README.md) | check for `TestCase` missing `CheckDestroy` | AST |
| [AT002](passes/AT002/README.md) | check for acceptance test function names including the word import | AST |
| [AT003](passes/AT003/README.md) | check for acceptance test function names missing an underscore | AST |
| [AT004](passes/AT004/README.md) | check for `TestStep` `Config` containing provider configuration | AST |
| [AT005](passes/AT005/README.md) | check for acceptance test function names missing `TestAcc` prefix | AST |
| [AT006](passes/AT006/README.md) | check for acceptance test functions containing multiple `resource.Test()` invocations | AST |
| [AT007](passes/AT007/README.md) | check for acceptance test functions containing multiple `resource.ParallelTest()` invocations | AST |
| [AT008](passes/AT008/README.md) | check for acceptance test function declaration `*testing.T` parameter naming | AST |

### Standard Resource Checks

| Check | Description | Type |
|---|---|---|
| [R001](passes/R001/README.md) | check for `ResourceData.Set()` calls using complex key argument | AST |
| [R002](passes/R002/README.md) | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| [R003](passes/R003/README.md) | check for `Resource` having `Exists` functions | AST |
| [R004](passes/R004/README.md) | check for `ResourceData.Set()` calls using incompatible value types | AST |
| [R005](passes/R005/README.md) | check for `ResourceData.HasChange()` calls that can be combined into one `HasChanges()` call | AST |
| [R006](passes/R006/README.md) | check for `RetryFunc` that omit retryable errors | AST |
| [R007](passes/R007/README.md) | check for deprecated `(schema.ResourceData).Partial` usage | AST |
| [R008](passes/R008/README.md) | check for deprecated `(schema.ResourceData).SetPartial` usage | AST |
| [R009](passes/R009/README.md) | check for Go panic usage | AST |
| [R010](passes/R010/README.md) | check for `(schema.ResourceData).GetChange` assignment which should use `(schema.ResourceData).Get` | AST |
| [R011](passes/R011/README.md) | check for `Resource` that configure `MigrateState` | AST |
| [R012](passes/R012/README.md) | check for data source `Resource` that configure `CustomizeDiff` | AST |
| [R013](passes/R013/README.md) | check for `map[string]*Resource` that resource names contain at least one underscore | AST |
| [R014](passes/R014/README.md) | check for `CreateFunc`, `DeleteFunc`, `ReadFunc`, and `UpdateFunc` parameter naming | AST |

### Standard Schema Checks

| Check | Description | Type |
|---|---|---|
| [S001](passes/S001/README.md) | check for `Schema` of `TypeList` or `TypeSet` missing `Elem` | AST |
| [S002](passes/S002/README.md) | check for `Schema` with both `Required` and `Optional` enabled | AST |
| [S003](passes/S003/README.md) | check for `Schema` with both `Required` and `Computed` enabled | AST |
| [S004](passes/S004/README.md) | check for `Schema` with both `Required` and `Default` configured | AST |
| [S005](passes/S005/README.md) | check for `Schema` with both `Computed` and `Default` configured | AST |
| [S006](passes/S006/README.md) | check for `Schema` of `TypeMap` missing `Elem` | AST |
| [S007](passes/S007/README.md) | check for `Schema` with both `Required` and `ConflictsWith` configured | AST |
| [S008](passes/S008/README.md) | check for `Schema` of `TypeList` or `TypeSet` with `Default` configured | AST |
| [S009](passes/S009/README.md) | check for `Schema` of `TypeList` or `TypeSet` with `ValidateFunc` configured | AST |
| [S010](passes/S010/README.md) | check for `Schema` of `Computed` only with `ValidateFunc` configured | AST |
| [S011](passes/S011/README.md) | check for `Schema` of `Computed` only with `DiffSuppressFunc` configured | AST |
| [S012](passes/S012/README.md) | check for `Schema` that `Type` is configured | AST |
| [S013](passes/S013/README.md) | check for `map[string]*Schema` that one of `Computed`, `Optional`, or `Required` is configured | AST |
| [S014](passes/S014/README.md) | check for `Schema` within `Elem` that `Computed`, `Optional`, and `Required` are not configured | AST |
| [S015](passes/S015/README.md) | check for `map[string]*Schema` that attribute names are valid | AST |
| [S016](passes/S016/README.md) | check for `Schema` that `Set` is only configured for `TypeSet` | AST |
| [S017](passes/S017/README.md) | check for `Schema` that `MaxItems` and `MinItems` are only configured for `TypeList`, `TypeMap`, or `TypeSet` | AST |
| [S018](passes/S018/README.md) | check for `Schema` that should use `TypeList` with `MaxItems: 1` | AST |
| [S019](passes/S019/README.md) | check for `Schema` that should omit `Computed`, `Optional`, or `Required` set to `false` | AST |
| [S020](passes/S020/README.md) | check for `Schema` of `Computed` only with `ForceNew` enabled | AST |
| [S021](passes/S021/README.md) | check for `Schema` that should omit `ComputedWhen` | AST |
| [S022](passes/S022/README.md) | check for `Schema` of `TypeMap` with invalid `Elem` of `*schema.Resource` | AST |
| [S023](passes/S023/README.md) | check for `Schema` that should omit `Elem` with incompatible `Type` | AST |
| [S024](passes/S024/README.md) | check for `Schema` that should omit `ForceNew` in data source schema attributes | AST |
| [S025](passes/S025/README.md) | check for `Schema` of `Computed` only with `AtLeastOneOf` configured | AST |
| [S026](passes/S026/README.md) | check for `Schema` of `Computed` only with `ConflictsWith` configured | AST |
| [S027](passes/S027/README.md) | check for `Schema` of `Computed` only with `Default` configured | AST |
| [S028](passes/S028/README.md) | check for `Schema` of `Computed` only with `DefaultFunc` configured | AST |
| [S029](passes/S029/README.md) | check for `Schema` of `Computed` only with `ExactlyOneOf` configured | AST |
| [S030](passes/S030/README.md) | check for `Schema` of `Computed` only with `InputDefault` configured | AST |
| [S031](passes/S031/README.md) | check for `Schema` of `Computed` only with `MaxItems` configured | AST |
| [S032](passes/S032/README.md) | check for `Schema` of `Computed` only with `MinItems` configured | AST |
| [S033](passes/S033/README.md) | check for `Schema` of `Computed` only with `StateFunc` configured | AST |
| [S034](passes/S034/README.md) | check for `Schema` that configure `PromoteSingle` | AST |
| [S035](passes/S035/README.md) | check for `Schema` with invalid `AtLeastOneOf` attribute references | AST |
| [S036](passes/S036/README.md) | check for `Schema` with invalid `ConflictsWith` attribute references | AST |
| [S037](passes/S037/README.md) | check for `Schema` with invalid `ExactlyOneOf` attribute references | AST |

### Standard Validation Checks

| Check | Description | Type |
|---|---|---|
| [V001](passes/V001/README.md) | check for custom `SchemaValidateFunc` that implement `validation.StringMatch()` or `validation.StringDoesNotMatch()` | AST |
| [V002](passes/V002/README.md) | check for deprecated `CIDRNetwork` validation function usage | AST |
| [V003](passes/V003/README.md) | check for deprecated `IPRange` validation function usage | AST |
| [V004](passes/V004/README.md) | check for deprecated `SingleIP` validation function usage | AST |
| [V005](passes/V005/README.md) | check for deprecated `ValidateJsonString` validation function usage | AST |
| [V006](passes/V006/README.md) | check for deprecated `ValidateListUniqueStrings` validation function usage | AST |
| [V007](passes/V007/README.md) | check for deprecated `ValidateRegexp` validation function usage | AST |
| [V008](passes/V008/README.md) | check for deprecated `ValidateRFC3339TimeString` validation function usage | AST |

## Extra Lint Checks

Extra lint checks are not included in the `tfproviderlint` tool and must be accessed via the `tfproviderlintx` tool or [added to a custom lint tool](#implementing-a-custom-lint-tool). Generally these represent advanced Terraform Plugin SDK functionality that is not appropriate for all Terraform Providers.

### Extra Resource Checks

| Check | Description | Type |
|---|---|---|
| [XR001](xpasses/XR001/README.md) | check for usage of `ResourceData.GetOkExists()` calls | AST |
| [XR002](xpasses/XR002/README.md) | check for `Resource` that should implement `Importer` | AST |
| [XR003](xpasses/XR003/README.md) | check for `Resource` that should implement `Timeouts` | AST |
| [XR004](xpasses/XR004/README.md) | check for `ResourceData.Set()` calls that should implement error checking with complex values | AST |

### Extra Schema Checks

| Check | Description | Type |
|---|---|---|
| [XS001](xpasses/XS001/README.md) | check for `map[string]*Schema` that `Description` is configured | AST |

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

The `helper/analysisfixtest` package contains functionality to verify `SuggestedFixes` by copying a source directory, running the `-fix` command against those copied files, then comparing expected file contents against those fixed files. Similar to the `go/analysis/analysistest` package, the testing can be invoked via:

```go
import (
  "testing"
  
  "github.com/bflad/tfproviderlint/helper/analysisfixtest"
  "golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerFixes(t *testing.T) {
  testdata := analysistest.TestData()
  analysisfixtest.Run(t, testdata, Analyzer, "a")
}
```

To setup the necessary expected file content verification, the testing expects a directory suffixed with `_fixed` (e.g. `testdata/src/a_fixed` for `testdata/src/a` files). The verification will only check files present in the `_fixed` directory against those fixed.

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
