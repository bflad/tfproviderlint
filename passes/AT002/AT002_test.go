package AT002_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/AT002"
)

func TestAT002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT002.Analyzer, "a")
}
