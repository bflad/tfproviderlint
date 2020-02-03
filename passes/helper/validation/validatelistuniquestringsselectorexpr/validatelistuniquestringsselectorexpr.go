package validatelistuniquestrings

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatelistuniquestrings",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateListUniqueStrings,
)
