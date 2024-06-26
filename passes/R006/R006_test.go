package R006_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/R006"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR006(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R006.Analyzer, "testdata/src/a")
}

func TestR006PackageAliases(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := R006.Analyzer
	analyzer.Flags.Set("package-aliases", "testdata/src/a/methodexpression")
	analysistest.Run(t, testdata, analyzer, "testdata/src/packagealiases")
}
