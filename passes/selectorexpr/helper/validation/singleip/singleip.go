package singleip

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"singleip",
	validation.IsPackageFunc,
	validation.PackagePath,
	validation.FuncNameSingleIP,
)
