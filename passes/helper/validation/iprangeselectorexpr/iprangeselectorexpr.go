package iprangeselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"iprange",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
