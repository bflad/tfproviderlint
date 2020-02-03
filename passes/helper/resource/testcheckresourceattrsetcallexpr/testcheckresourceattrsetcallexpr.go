package testcheckresourceattrsetcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattrsetcallexpr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttrSet,
)
