package providertypeassertexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.TypeAssertExprAnalyzer(
	"providertypeassertexpr",
	schema.IsFunc,
	schema.PackagePath,
	schema.TypeNameProvider,
)
