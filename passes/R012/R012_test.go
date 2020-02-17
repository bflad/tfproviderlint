package R012

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestR012(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
