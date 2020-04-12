package S038

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.SchemaAttributeReferencesSemanticsAnalyzer("S038", schema.SchemaFieldAtLeastOneOf)
