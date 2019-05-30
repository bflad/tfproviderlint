package S004_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S004"
)

func TestS004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S004.Analyzer, "a")
}
