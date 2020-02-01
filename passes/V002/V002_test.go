package V002_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/V002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V002.Analyzer, "a")
}
