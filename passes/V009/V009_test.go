package V009_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/V009"
)

func TestV009(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V009.Analyzer, "a")
}
