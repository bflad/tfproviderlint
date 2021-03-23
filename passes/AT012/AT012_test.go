package AT012_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/AT012"
)

func TestAT012(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT012.Analyzer, "a")
}

func TestAT012IgnoredFilenames(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := AT012.Analyzer
	analyzer.Flags.Set("ignored-filenames", "ignored_file_test.go")
	analysistest.Run(t, testdata, analyzer, "ignoredfilenames")
}
