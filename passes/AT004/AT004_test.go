package AT004_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/AT004"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT004.Analyzer, "a")
}
