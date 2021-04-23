package XS003

import (
	"go/ast"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Schema that might introduce crash or diff as it allows to be an empty block

The XS003 analyzer reports cases of schemas where a non-computed nested block property of
"TypeList" contains optional-only and non-default child properties, which has no "AtLeastOneOf"/"ExactlyOneOf" constraint.`

const analyzerName = "XS003"

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

schemaLoop:
	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		// Ignore if this is not a non-computed nested block of TypeList.
		if (!schemaInfo.Schema.Optional && !schemaInfo.Schema.Required) ||
			schemaInfo.SchemaValueType != schema.SchemaValueTypeList || schemaInfo.Schema.Elem == nil {
			continue
		}

		elem := schemaInfo.Schema.Elem
		resource, ok := elem.(*schema.ResourceType)
		if !ok || resource == nil {
			continue
		}

		// Ignore if the child properties of this nested block has at least one required property, or Default/DefaultFunc,
		// or one of "AtLeastOnOf"/"ExactlyOneOf" constraint.
		for _, prop := range resource.Schema {
			if prop == nil {
				continue schemaLoop
			}
			if prop.Required || prop.Default != nil || prop.DefaultFunc != nil || prop.AtLeastOneOf != nil || prop.ExactlyOneOf != nil {
				continue schemaLoop
			}
		}

		switch t := schemaInfo.AstCompositeLit.Type.(type) {
		default:
			pass.Reportf(schemaInfo.AstCompositeLit.Lbrace, "%s: schema might introduce crash or diff as it allows to be an empty block", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema might introduce crash or diff as it allows to be an empty block", analyzerName)
		}
	}

	return nil, nil
}
