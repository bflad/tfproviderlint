package validateregexp

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validateregexp",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
)
