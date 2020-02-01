package testmatchresourceattr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testmatchresourceattr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestMatchResourceAttr,
)
