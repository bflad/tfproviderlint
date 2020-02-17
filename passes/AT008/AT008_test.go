package AT008

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT008(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
