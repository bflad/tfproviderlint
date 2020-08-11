package XS002_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XS002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXS002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XS002.Analyzer, "a")
}
