package XR008_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XR008"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR008.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	t.Skipf("skipping as these suggested fixes will intentionally create an invalid file")

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, XR008.Analyzer, "a")
}
