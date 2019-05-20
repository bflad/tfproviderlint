package resourcedatasetvaluederef_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/resourcedatasetvaluederef"
)

func TestResourceDataSetValueDeref(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, resourcedatasetvaluederef.Analyzer, "a")
}
