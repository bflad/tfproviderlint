package validatejsonstring

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstring",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
)
