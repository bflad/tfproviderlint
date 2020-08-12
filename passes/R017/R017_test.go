package R017_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R017"
)

func TestR017(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R017.Analyzer, "a")
}
