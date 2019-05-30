package R003_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R003"
)

func TestR003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R003.Analyzer, "a")
}
