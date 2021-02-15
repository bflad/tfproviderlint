package XR006_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XR006"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR006(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR006.Analyzer, "a")
}
