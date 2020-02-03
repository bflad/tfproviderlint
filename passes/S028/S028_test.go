package S028_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S028"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS028(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S028.Analyzer, "a")
}
