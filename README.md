# tfproviderlint

A linting tool for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

## Install

### Local Install

Release binaries are available in the [Releases](https://github.com/bflad/tfproviderlint/releases) section.

To instead use Go to install into your `$GOBIN` directory (e.g. `$GOPATH/bin`):

```console
$ go get github.com/bflad/tfproviderlint/cmd/tfproviderlint
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

Additional information about usage and configuration options can be found by passing the `help` argument:

```console
$ tfproviderlint help
```

### Local Usage

Change into the directory of the Terraform Provider code and run:

```console
$ tfproviderlint ./...
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

## Lint Checks

For additional information about each check, you can run `tfproviderlint help NAME`.

### Acceptance Test Checks

| Check | Description | Type |
|---|---|---|
| [AT001](passes/AT001/README.md) | check for `TestCase` missing `CheckDestroy` | AST |
| AT002 | check for acceptance test function names including the word import | AST |
| AT003 | check for acceptance test function names missing an underscore | AST |
| AT004 | check for `TestStep` `Config` containing provider configuration | AST |

### Resource Checks

| Check | Description | Type |
|---|---|---|
| R001 | check for `ResourceData.Set()` calls using complex key argument | AST |
| R002 | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| R003 | check for `Resource` having `Exists` functions | AST |
| R004 | check for `ResourceData.Set()` calls using incompatible value types | AST |

### Schema Checks

| Check | Description | Type |
|---|---|---|
| S001 | check for `Schema` of `TypeList` or `TypeSet` missing `Elem` | AST |
| S002 | check for `Schema` with both `Required` and `Optional` enabled | AST |
| S003 | check for `Schema` with both `Required` and `Computed` enabled | AST |
| S004 | check for `Schema` with both `Required` and `Default` configured | AST |
| S005 | check for `Schema` with both `Computed` and `Default` configured | AST |
| S006 | check for `Schema` of `TypeMap` missing `Elem` | AST |

## Development and Testing

This project is built on the [`go/analysis`](https://godoc.org/golang.org/x/tools/go/analysis) framework and uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

Helpful tooling for development:

* [`astdump`](https://github.com/wingyplus/astdump): a tool for displaying the AST form of Go file
* [`ssadump`](https://godoc.org/golang.org/x/tools/cmd/ssadump): a tool for displaying and interpreting the SSA form of Go programs

### Adding an Analyzer

* Create new analyzer in `passes/`
* Add new analyzer in `cmd/tfproviderlint/tfproviderlint.go`
* Since the [`analysistest` package](https://godoc.org/golang.org/x/tools/go/analysis/analysistest) does not support Go Modules currently, each analyzer that implements testing must add a symlink to the top level `vendor` directory in the `testdata/src` directory. e.g. `ln -s ../../../../vendor passes/NAME/testdata/src/vendor`

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
