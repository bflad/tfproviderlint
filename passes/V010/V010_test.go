package V010_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/V010"
)

func TestV010(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V010.Analyzer, "a")
}
