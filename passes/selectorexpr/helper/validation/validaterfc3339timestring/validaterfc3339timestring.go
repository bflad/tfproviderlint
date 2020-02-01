package validaterfc3339timestring

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validaterfc3339timestring",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameValidateRFC3339TimeString,
)
