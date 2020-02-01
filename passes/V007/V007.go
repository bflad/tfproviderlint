package V007

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/validateregexp"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V007",
	validateregexp.Analyzer,
	validation.PackageName,
	validation.FuncNameValidateRegexp,
	validation.PackageName,
	validation.FuncNameStringIsValidRegExp,
)
