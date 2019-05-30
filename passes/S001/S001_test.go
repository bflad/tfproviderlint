package S001_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S001"
)

func TestS001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S001.Analyzer, "a")
}
