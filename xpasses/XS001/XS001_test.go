package XS001_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XS001"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXS001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XS001.Analyzer, "a")
}
