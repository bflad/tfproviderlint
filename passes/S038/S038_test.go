package S038

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS038(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
