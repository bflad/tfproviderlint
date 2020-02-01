package V002

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/cidrnetwork"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V002",
	cidrnetwork.Analyzer,
	validation.PackageName,
	validation.FuncNameCIDRNetwork,
	validation.PackageName,
	validation.FuncNameIsCIDRNetwork,
)
