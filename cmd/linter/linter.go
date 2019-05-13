// The linter command is a static checker for Terraform Providers.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestcheckdestroy"
	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestproviderconfig"
)

func main() {
	multichecker.Main(
		acctestcheckdestroy.Analyzer,
		acctestproviderconfig.Analyzer,
	)
}
