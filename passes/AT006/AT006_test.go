package AT006_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/AT006"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT006(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT006.Analyzer, "a")
}
