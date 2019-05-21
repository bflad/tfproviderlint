package acctestproviderconfig_test

import (
	"testing"

	"github.com/bflad/tfproviderlint/passes/acctestproviderconfig"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAccTestProviderConfig(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestproviderconfig.Analyzer, "a")
}
