package cidrnetwork

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"cidrnetwork",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
)
