package iprange

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"iprange",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
