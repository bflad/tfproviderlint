package schema

import (
	"go/ast"
	"reflect"
	"strings"
	"testing"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema/testdata/type_resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.org/x/tools/go/packages"
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
		desc := ast.NewCommentMap(pkg.Fset, resourceFuncLit, f.Comments)[resourceFuncLit][i].List[0].Text

		var expectResource *schema.Resource
		if strings.Contains(desc, "no panic") {
			expectResource = &schema.Resource{
				Schema: map[string]*schema.Schema{},
			}
		} else {
			expectResource = type_resource.ResourceFuncs[i]()
		}

		// new resource info
		cl := resourceFuncLit.(*ast.FuncLit).Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit)
		resourceInfo := NewResourceInfo(cl, pkg.TypesInfo)

		// test
		if !reflect.DeepEqual(*expectResource, *resourceInfo.Resource) {
			t.Fatalf("test case: %s, resource not match\n\nexpected: %#v\n\ngot: %#v\n\n", desc, *expectResource, *resourceInfo.Resource)
		}
	}
}
