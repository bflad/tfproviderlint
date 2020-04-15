package analysisutils

import (
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"
)

// SchemaAttributeReferencesSyntaxRunner returns an Analyzer runner for fields that use schema attribute references
func SchemaAttributeReferencesSyntaxRunner(analyzerName string, fieldName string) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
		schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)

		for _, schemaInfo := range schemaInfos {
			if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
				continue
			}

			if !schemaInfo.DeclaresField(fieldName) {
				continue
			}

			switch value := schemaInfo.Fields[fieldName].Value.(type) {
			case *ast.CompositeLit:
				if !astutils.IsExprTypeArrayString(value.Type) {
					continue
				}

				for _, elt := range value.Elts {
					attributeReference := astutils.ExprStringValue(elt)

					if attributeReference == nil {
						continue
					}

					if _, err := schema.ParseAttributeReference(*attributeReference); err != nil {
						pass.Reportf(elt.Pos(), "%s: invalid %s attribute reference syntax: %s", analyzerName, fieldName, err)
					}
				}
			}
		}

		return nil, nil
	}
}

// SchemaAttributeReferencesSemanticsRunner returns an Analyzer runner for fields that use schema attribute references
func SchemaAttributeReferencesSemanticsRunner(analyzerName string, fieldName string) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
		resourceInfos := pass.ResultOf[resourceinfo.Analyzer].([]*schema.ResourceInfo)

		for _, resourceInfo := range resourceInfos  {
			// Handle "root" schema.Resource only since we will use the resourceInfo of that to validate attribute reference.
			if !resourceInfo.IsDataSource() && !resourceInfo.IsResource() {
				continue
			}

			// Collection schemaInfo under current resourceInfo
			var schemaInfos []*schema.SchemaInfo
			ast.Inspect(resourceInfo.AstCompositeLit, func(n ast.Node) bool {
				compositeLit, ok := n.(*ast.CompositeLit)

				if !ok {
					return true
				}

				if schema.IsMapStringSchema(compositeLit, pass.TypesInfo) {
					for _, mapSchema := range schema.GetSchemaMapSchemas(compositeLit) {
						schemaInfos = append(schemaInfos, schema.NewSchemaInfo(mapSchema, pass.TypesInfo))
					}
				} else if schema.IsTypeSchema(pass.TypesInfo.TypeOf(compositeLit.Type)) {
					schemaInfos = append(schemaInfos, schema.NewSchemaInfo(compositeLit, pass.TypesInfo))
				}

				return true
			})

			// Validate each schemaInfo under current resourceInfo
			for _, schemaInfo := range schemaInfos {
				if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
					continue
				}

				if !schemaInfo.DeclaresField(fieldName) {
					continue
				}

				switch value := schemaInfo.Fields[fieldName].Value.(type) {
				case *ast.CompositeLit:
					if !astutils.IsExprTypeArrayString(value.Type) {
						continue
					}

					for _, elt := range value.Elts {
						attributeReference := astutils.ExprStringValue(elt)

						if attributeReference == nil {
							continue
						}

						if err := schema.ValidateAttributeReference(resourceInfo.Resource, *attributeReference); err != nil {
							pass.Reportf(elt.Pos(), "%s: invalid %s attribute reference semantics: %s", analyzerName, fieldName, err)
						}
					}
				}
			}
		}

		return nil, nil
	}
}
