package S014_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S014"
)

func TestS014(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S014.Analyzer, "a")
}
