// The linter command is a static checker for Terraform Providers.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/bflad/tfproviderlint/passes/acctestcheckdestroy"
	"github.com/bflad/tfproviderlint/passes/acctestfuncnameimport"
	"github.com/bflad/tfproviderlint/passes/acctestfuncnameunderscore"
	"github.com/bflad/tfproviderlint/passes/acctestproviderconfig"
	"github.com/bflad/tfproviderlint/passes/resourcedatasetkey"
	"github.com/bflad/tfproviderlint/passes/resourcedatasetvaluederef"
	"github.com/bflad/tfproviderlint/passes/resourceexistsfunc"
	"github.com/bflad/tfproviderlint/passes/schematypemapelem"
)

func main() {
	multichecker.Main(
		acctestcheckdestroy.Analyzer,
		acctestfuncnameimport.Analyzer,
		acctestfuncnameunderscore.Analyzer,
		acctestproviderconfig.Analyzer,
		resourcedatasetkey.Analyzer,
		resourcedatasetvaluederef.Analyzer,
		resourceexistsfunc.Analyzer,
		schematypemapelem.Analyzer,
	)
}
