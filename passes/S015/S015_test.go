package S015_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S015"
)

func TestS015(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S015.Analyzer, "a")
}
