package S037

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS037(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
