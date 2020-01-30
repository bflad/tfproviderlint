package XR002_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XR002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR002.Analyzer, "a")
}
