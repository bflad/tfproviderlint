package testcheckresourceattrcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttr,
)
