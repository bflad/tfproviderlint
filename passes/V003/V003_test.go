package V003_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/helper/analysisfixtest"
	"github.com/bflad/tfproviderlint/passes/V003"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V003.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysisfixtest.Run(t, testdata, V003.Analyzer, "a")
}
