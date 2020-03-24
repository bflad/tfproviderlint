package schema

import (
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema/testdata/type_schema"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"reflect"
	"testing"
)


func TestNewSchemaInfo_Schema(t *testing.T) {
	// load the single test file
	const testfile = "testdata/type_schema/schema.go"
	cfg := &packages.Config{Mode: packages.LoadAllSyntax}
	pkgs, _ := packages.Load(cfg, testfile)
	pkg := pkgs[0]
	f := pkg.Syntax[0]

	schemaFuncLits := f.Decls[2].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CompositeLit).Elts

	for i, schemaFuncLit := range schemaFuncLits {
		expectSchema := type_schema.SchemaFuncs[i]()
		desc := ast.NewCommentMap(pkg.Fset, schemaFuncLit, f.Comments)[schemaFuncLit][0].List[0].Text

		// new schema info
		cl := schemaFuncLit.(*ast.FuncLit).Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit)
		schemaInfo := NewSchemaInfo(cl, pkg.TypesInfo)

		// test
		if !reflect.DeepEqual(expectSchema, schemaInfo.Schema) {
			t.Fatalf("test case: %s, schema not match", desc)
		}
	}
}
