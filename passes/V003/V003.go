package V003

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/iprange"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V003",
	iprange.Analyzer,
	validation.PackageName,
	validation.FuncNameIPRange,
	validation.PackageName,
	validation.FuncNameIsIPv4Range,
)
