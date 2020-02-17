package R011

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR011(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
