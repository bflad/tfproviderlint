package R014

import (
	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/crudfuncinfo"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for CreateFunc, DeleteFunc, ReadFunc, and UpdateFunc parameter naming

The R014 analyzer reports when CreateFunc, DeleteFunc, ReadFunc, and UpdateFunc
declarations do not use d as the name for the *schema.ResourceData parameter
or meta as the name for the interface{} parameter. This parameter naming is the
standard convention for resources.`

const analyzerName = "R014"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		crudfuncinfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	crudFuncs := pass.ResultOf[crudfuncinfo.Analyzer].([]*schema.CRUDFuncInfo)

	for _, crudFunc := range crudFuncs {
		if ignorer.ShouldIgnore(analyzerName, crudFunc.Node) {
			continue
		}

		params := crudFunc.Type.Params
		paramCount := len(params.List)

		if name := astutils.FieldListName(params, paramCount-2, 0); name != nil && *name != "_" && *name != "d" {
			pass.Reportf(params.List[paramCount-2].Pos(), "%s: *schema.ResourceData parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named d", analyzerName)
		}

		if name := astutils.FieldListName(params, paramCount-1, 0); name != nil && *name != "_" && *name != "meta" {
			pass.Reportf(params.List[paramCount-1].Pos(), "%s: interface{} parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named meta", analyzerName)
		}
	}

	return nil, nil
}
