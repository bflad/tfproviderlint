package AT011_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/AT011"
)

func TestAT011(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT011.Analyzer, "a")
}
