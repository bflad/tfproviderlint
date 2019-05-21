package schematypemapelem_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/bflad/tfproviderlint/passes/schematypemapelem"
)

func TestSchemaTypeMapElem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, schematypemapelem.Analyzer, "a")
}
