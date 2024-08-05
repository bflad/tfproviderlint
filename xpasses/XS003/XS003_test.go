package XS003_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/xpasses/XS003"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXS003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XS003.Analyzer, "a")
}
