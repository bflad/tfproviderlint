package resourceproviderfactoryselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderfactoryselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProviderFactory,
)
