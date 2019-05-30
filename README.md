# tfproviderlint

A linting tool for [Terraform Provider](https://www.terraform.io/docs/providers/index.html) code.

## Install

### Local Install

To install the linter in your `$GOBIN` directory (e.g. `$GOPATH/bin`):

```console
$ go get github.com/bflad/tfproviderlint/cmd/tfproviderlint
```

### Docker Install

```console
$ docker pull bflad/tfproviderlint
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
| AT001 | check for `TestCase` missing `CheckDestroy` | AST |
| AT002 | check for acceptance test function names including the word import | AST |
| AT003 | check for acceptance test function names missing an underscore | AST |
| AT004 | check for `TestStep` `Config` containing provider configuration | AST |

### Resource Checks

| Check | Description | Type |
|---|---|---|
| R001 | check for `ResourceData.Set()` calls using complex key argument | AST |
| R002 | check for `ResourceData.Set()` calls using `*` dereferences | AST |
| R003 | check for `Resource` having `Exists` functions | AST |

### Schema Checks

| Check | Description | Type |
|---|---|---|
| S001 | check for `Schema` of `TypeMap` missing `Elem` | AST |
| S002 | check for `Schema` with both `Required` and `Optional` enabled | AST |
| S003 | check for `Schema` with both `Required` and `Computed` enabled | AST |

## Development and Testing

This project is built on the [`go/analysis`](https://godoc.org/golang.org/x/tools/go/analysis) framework and uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

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

### Docker Build Testing

```console
$ docker build .
```

### Unit Testing

```console
$ go test ./...
```
