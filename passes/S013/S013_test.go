package S013_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S013"
)

func TestS013(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S013.Analyzer, "a")
}
