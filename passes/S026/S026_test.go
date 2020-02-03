package S026_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S026"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS026(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S026.Analyzer, "a")
}
