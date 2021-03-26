package XR007_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XR007"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR007.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, XR007.Analyzer, "a")
}
