package S016_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S016"
)

func TestS016(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S016.Analyzer, "testdata/src/a")
}
