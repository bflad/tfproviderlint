package XAT001_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/xpasses/XAT001"
)

func TestXAT001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XAT001.Analyzer, "a")
}
