package S032_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S032"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS032(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S032.Analyzer, "a")
}
