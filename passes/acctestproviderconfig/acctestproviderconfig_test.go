package acctestproviderconfig_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctestproviderconfig"
)

func TestAccTestProviderConfig(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctestproviderconfig.Analyzer, "a")
}
