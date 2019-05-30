package S002_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S002"
)

func TestS002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S002.Analyzer, "a")
}
