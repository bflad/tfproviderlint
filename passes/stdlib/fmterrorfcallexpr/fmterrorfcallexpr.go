package fmterrorfcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmterrorfcallexpr",
	"fmt",
	"Errorf",
)
