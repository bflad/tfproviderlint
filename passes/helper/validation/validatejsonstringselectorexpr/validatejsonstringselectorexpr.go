package validatejsonstring

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstring",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
)
