package AT001_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/AT001"
)

func TestAT001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT001.Analyzer, "a")
}

func TestAT001CustomSuffixes(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := AT001.Analyzer
	analyzer.Flags.Set("ignored-filename-suffixes", "_data_source_test.go")
	analysistest.Run(t, testdata, analyzer, "suffixes")
}

func TestAT001CustomPrefixes(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := AT001.Analyzer
	analyzer.Flags.Set("ignored-filename-prefixes", "extra_data_")
	analysistest.Run(t, testdata, analyzer, "prefixes")
}