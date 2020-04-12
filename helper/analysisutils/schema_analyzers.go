package analysisutils

import (
	"fmt"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"
)

// SchemaAttributeReferencesSyntaxAnalyzer returns an Analyzer for fields that use schema attribute references
func SchemaAttributeReferencesSyntaxAnalyzer(analyzerName string, fieldName string) *analysis.Analyzer {
	doc := fmt.Sprintf(`check for Schema %[2]s references with invalid syntax

The %[1]s analyzer ensures schema attribute references in the Schema %[2]s
field use valid syntax. The Terraform Plugin SDK can unit test attribute
references to verify the references against the full schema.
`, analyzerName, fieldName)

	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  doc,
		Requires: []*analysis.Analyzer{
			commentignore.Analyzer,
			schemainfo.Analyzer,
		},
		Run: SchemaAttributeReferencesSyntaxRunner(analyzerName, fieldName),
}
}

// SchemaAttributeReferencesSemanticsAnalyzer returns an Analyzer for fields that use schema attribute references
func SchemaAttributeReferencesSemanticsAnalyzer(analyzerName string, fieldName string) *analysis.Analyzer {
	doc := fmt.Sprintf(`check for Schema %[2]s references with invalid semantics

The %[1]s analyzer ensures schema attribute references in the Schema %[2]s
field refers to valid attribute.
`, analyzerName, fieldName)

	return &analysis.Analyzer{
		Name: analyzerName,
		Doc:  doc,
		Requires: []*analysis.Analyzer{
			commentignore.Analyzer,
			resourceinfo.Analyzer,
		},
		Run: SchemaAttributeReferencesSemanticsRunner(analyzerName, fieldName),
	}
}
