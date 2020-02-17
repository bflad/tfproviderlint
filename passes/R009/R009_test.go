package R009_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R009"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR009(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R009.Analyzer, "a")
}
