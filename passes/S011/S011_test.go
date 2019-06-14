package S011_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S011"
)

func TestS011(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S011.Analyzer, "a")
}
