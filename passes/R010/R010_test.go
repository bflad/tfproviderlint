package R010_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R010"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR010(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R010.Analyzer, "a")
}
