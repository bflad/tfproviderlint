package S009_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/S009"
)

func TestS009(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S009.Analyzer, "a")
}
