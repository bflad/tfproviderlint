package R019_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R019"
)

func TestR019(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := R019.Analyzer
	analysistest.Run(t, testdata, analyzer, "a")
}

func TestR019Threshold(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := R019.Analyzer

	if err := analyzer.Flags.Set("threshold", "3"); err != nil {
		t.Fatalf("error setting threshold flag: %s", err)
	}

	analysistest.Run(t, testdata, analyzer, "threshold")
}
