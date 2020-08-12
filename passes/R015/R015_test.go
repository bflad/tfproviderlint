package R015_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/R015"
)

func TestR015(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R015.Analyzer, "a")
}
