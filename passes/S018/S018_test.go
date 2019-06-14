package S018_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S018"
)

func TestS018(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S018.Analyzer, "a")
}
