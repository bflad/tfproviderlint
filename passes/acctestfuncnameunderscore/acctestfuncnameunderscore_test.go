package acctestfuncnameunderscore_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/acctestfuncnameunderscore"
)

func TestAccTestFuncNameUnderscore(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestfuncnameunderscore.Analyzer, "a")
}
