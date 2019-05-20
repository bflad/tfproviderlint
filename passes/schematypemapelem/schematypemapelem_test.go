package schematypemapelem_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/terraform-providers/terraform-provider-aws/linter/passes/schematypemapelem"
)

func TestSchemaTypeMapElem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, schematypemapelem.Analyzer, "a")
}
