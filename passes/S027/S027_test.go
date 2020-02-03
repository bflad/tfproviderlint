package S027_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S027"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS027(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S027.Analyzer, "a")
}
