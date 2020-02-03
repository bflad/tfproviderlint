package S031_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S031"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS031(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S031.Analyzer, "a")
}
