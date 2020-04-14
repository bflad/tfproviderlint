package R008_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R008"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R008.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, R008.Analyzer, "a")
}
