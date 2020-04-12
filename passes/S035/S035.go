package S035

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.SchemaAttributeReferencesSyntaxAnalyzer("S035", schema.SchemaFieldAtLeastOneOf)
