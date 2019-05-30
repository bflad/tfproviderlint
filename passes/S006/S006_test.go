package S006_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S006"
)

func TestS006(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S006.Analyzer, "a")
}
