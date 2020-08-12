package R016_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R016"
)

func TestR016(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R016.Analyzer, "a")
}
