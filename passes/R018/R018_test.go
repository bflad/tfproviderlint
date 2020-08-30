package R018_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R018"
)

func TestR018(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R018.Analyzer, "a")
}
