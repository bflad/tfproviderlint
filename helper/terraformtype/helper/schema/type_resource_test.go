package schema

import (
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema/testdata/type_resource"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"reflect"
	"testing"
)


func TestNewResourceInfo_Resource(t *testing.T) {
	// load the single test file
	const testfile = "testdata/type_resource/resource.go"
	cfg := &packages.Config{Mode: packages.LoadAllSyntax}
	pkgs, _ := packages.Load(cfg, testfile)
	pkg := pkgs[0]
	f := pkg.Syntax[0]

	resourceFuncLits := f.Decls[2].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CompositeLit).Elts

	for i, resourceFuncLit := range resourceFuncLits {
		expectResource := type_resource.ResourceFuncs[i]()
		desc := ast.NewCommentMap(pkg.Fset, resourceFuncLit, f.Comments)[resourceFuncLit][0].List[0].Text

		// new resource info
		cl := resourceFuncLit.(*ast.FuncLit).Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit)
		resourceInfo := NewResourceInfo(cl, pkg.TypesInfo)

		// test
		if !reflect.DeepEqual(expectResource, resourceInfo.Resource) {
			t.Fatalf("test case: %s, resource not match", desc)
		}
	}
}
