package S007_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S007"
)

func TestS007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S007.Analyzer, "a")
}
