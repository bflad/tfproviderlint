package R013

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR013(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
