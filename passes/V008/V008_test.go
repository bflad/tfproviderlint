package V008_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/helper/analysisfixtest"
	"github.com/bflad/tfproviderlint/passes/V008"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V008.Analyzer, "a")
}

func TestAnalyzerFixes(t *testing.T) {
	testdata := analysistest.TestData()
	analysisfixtest.Run(t, testdata, V008.Analyzer, "a")
}
