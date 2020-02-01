package V004

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/singleip"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V004",
	singleip.Analyzer,
	validation.PackageName,
	validation.FuncNameSingleIP,
	validation.PackageName,
	validation.FuncNameIsIPAddress,
)
