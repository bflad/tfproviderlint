package R014

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR014(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
