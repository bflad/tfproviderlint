package V001_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/V001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V001.Analyzer, "a")
}
