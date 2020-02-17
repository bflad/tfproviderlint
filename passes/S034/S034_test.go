package S034

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS034(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
