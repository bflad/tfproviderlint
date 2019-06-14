package S017_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S017"
)

func TestS017(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S017.Analyzer, "a")
}
