// Package S038 defines an Analyzer that checks for
// Schema that should omit Elem with incompatible Type
package S038

import (
	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema that has Default value declared with incompatible Type

The S038 analyzer reports cases of schema that declare Default value with incompatible Type.`

const analyzerName = "S038"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemainfo.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)
	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		// If there is no "Default" declared in the schema, continue
		if kvExpr := schemaInfo.Fields[schema.SchemaFieldDefault]; kvExpr == nil || astutils.ExprValue(kvExpr.Value) == nil {
			continue
		}

		// "Default" is declared with an non-nil value. Whilst, its value doesn't match the schema type.
		// (the .Schema.Default will be correctly set during NewSchemaInfo() only if the type matches)
		if schemaInfo.Schema.Default != nil {
			continue
		}

		pass.Reportf(schemaInfo.Fields[schema.SchemaFieldDefault].Value.Pos(), "%s: schema should not declare Default with incompatible Type", analyzerName)
	}

	return nil, nil
}
