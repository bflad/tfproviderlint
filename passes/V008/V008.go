package V008

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/validaterfc3339timestring"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V008",
	validaterfc3339timestring.Analyzer,
	validation.PackageName,
	validation.FuncNameValidateRFC3339TimeString,
	validation.PackageName,
	validation.FuncNameIsRFC3339Time,
)
