package validateregexpselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validateregexp",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
)
