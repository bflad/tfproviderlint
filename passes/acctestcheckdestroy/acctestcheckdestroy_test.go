package acctestcheckdestroy_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/acctestcheckdestroy"
)

func TestAccTestCheckDestroy(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestcheckdestroy.Analyzer, "a")
}
