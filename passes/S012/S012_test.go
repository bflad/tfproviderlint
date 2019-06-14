package S012_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S012"
)

func TestS012(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S012.Analyzer, "a")
}
