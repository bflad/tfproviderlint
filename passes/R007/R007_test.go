package R007_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R007"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R007.Analyzer, "a")
}
