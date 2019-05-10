// The vet command is a static checker for Go programs. It has pluggable
// analyzers defined using the golang.org/x/tools/go/analysis API, and
// using the golang.org/x/tools/go/packages API to load packages in any
// build system.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestcheckdestroy"
)

func main() {
	multichecker.Main(
		acctestcheckdestroy.Analyzer,
	)
}
