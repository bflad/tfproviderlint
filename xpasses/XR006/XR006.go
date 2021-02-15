package XR006

import (
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
)

const Doc = `check for Resource that implements redundant Timeouts

The XR006 analyzer reports redundant Timeouts in resources.`

const analyzerName = "XR006"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		resourceinfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resources := pass.ResultOf[resourceinfo.Analyzer].([]*schema.ResourceInfo)
	for _, resource := range resources {
		if ignorer.ShouldIgnore(analyzerName, resource.AstCompositeLit) {
			continue
		}

		if (resource.Fields[schema.ResourceFieldCreate] == nil ) == (resource.Resource.Timeout.Create == nil ) &&
			(resource.Fields[schema.ResourceFieldRead] == nil ) == (resource.Resource.Timeout.Read == nil ) &&
			(resource.Fields[schema.ResourceFieldUpdate] == nil ) == (resource.Resource.Timeout.Update == nil ) &&
			(resource.Fields[schema.ResourceFieldDelete] == nil ) == (resource.Resource.Timeout.Delete == nil ) {
			continue
		}


		pass.Reportf(resource.AstCompositeLit.Pos(), "%s: resource should not include redundant Timeouts implementation", analyzerName)
	}

	return nil, nil
}
