package S023_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/S023"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS023(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S023.Analyzer, "a")
}
