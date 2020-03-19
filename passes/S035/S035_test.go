package S035

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS035(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
