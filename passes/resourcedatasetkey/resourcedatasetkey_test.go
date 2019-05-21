package resourcedatasetkey_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/resourcedatasetkey"
)

func TestResourceDataSetKey(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, resourcedatasetkey.Analyzer, "a")
}
