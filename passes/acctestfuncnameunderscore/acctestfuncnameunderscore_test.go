package acctestfuncnameunderscore_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestfuncnameunderscore"
)

func TestAccTestFuncNameUnderscore(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestfuncnameunderscore.Analyzer, "a")
}
