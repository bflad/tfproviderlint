package resourceexistsfunc_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/resourceexistsfunc"
)

func TestResourceExistsFunc(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, resourceexistsfunc.Analyzer, "a")
}
