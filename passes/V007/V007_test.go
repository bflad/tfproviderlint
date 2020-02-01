package V007_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/V007"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V007.Analyzer, "a")
}
