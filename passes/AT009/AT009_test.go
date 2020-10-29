package AT009

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT009(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
