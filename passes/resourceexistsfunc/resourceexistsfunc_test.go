package resourceexistsfunc_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/resourceexistsfunc"
)

func TestResourceExistsFunc(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, resourceexistsfunc.Analyzer, "a")
}
