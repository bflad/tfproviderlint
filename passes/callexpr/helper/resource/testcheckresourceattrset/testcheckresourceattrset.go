package testcheckresourceattrset

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattrset",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttrSet,
)
