package V005

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/helper/validation/validatejsonstringselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V005",
	validatejsonstring.Analyzer,
	validation.PackageName,
	validation.FuncNameValidateJsonString,
	validation.PackageName,
	validation.FuncNameStringIsJSON,
)
