package S030_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S030"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS030(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S030.Analyzer, "a")
}
