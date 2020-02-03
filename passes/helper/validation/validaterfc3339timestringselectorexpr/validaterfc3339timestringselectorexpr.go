package validaterfc3339timestringselectorexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validaterfc3339timestring",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRFC3339TimeString,
)
