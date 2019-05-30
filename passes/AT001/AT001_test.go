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
