package AT003_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/AT003"
)

func TestAT003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT003.Analyzer, "a")
}
