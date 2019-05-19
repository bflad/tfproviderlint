package acctestfuncnameimport_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestfuncnameimport"
)

func TestAccTestFuncNameImport(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestfuncnameimport.Analyzer, "a")
}
