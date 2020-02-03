package S029_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S029"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS029(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S029.Analyzer, "a")
}
