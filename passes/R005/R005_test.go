package R005_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R005"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R005.Analyzer, "a")
}
