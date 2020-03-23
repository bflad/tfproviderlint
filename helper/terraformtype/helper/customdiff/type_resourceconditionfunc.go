package customdiff

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

const (
	TypeNameResourceConditionFunc = `ResourceConditionFunc`
)

// IsFuncTypeResourceConditionFunc returns true if the FuncType matches expected parameters and results types
func IsFuncTypeResourceConditionFunc(node ast.Node, info *types.Info) bool {
	funcType := astutils.FuncTypeFromNode(node)

	if funcType == nil {
		return false
	}

	if !astutils.HasFieldListLength(funcType.Params, 2) {
		return false
	}

	if !astutils.IsFieldListTypePackageType(funcType.Params, 0, info, schema.PackagePath, schema.TypeNameResourceDiff) {
		return false
	}

	if !astutils.IsFieldListType(funcType.Params, 1, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.HasFieldListLength(funcType.Results, 1) {
		return false
	}

	return astutils.IsFieldListType(funcType.Results, 0, astutils.IsExprTypeBool)
}

// IsTypeResourceConditionFunc returns if the type is ResourceConditionFunc from the customdiff package
func IsTypeResourceConditionFunc(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameResourceConditionFunc)
	case *types.Pointer:
		return IsTypeResourceConditionFunc(t.Elem())
	default:
		return false
	}
}

// ResourceConditionFuncInfo represents all gathered ResourceConditionFunc data for easier access
type ResourceConditionFuncInfo struct {
	AstFuncDecl *ast.FuncDecl
	AstFuncLit  *ast.FuncLit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.FuncType
	TypesInfo   *types.Info
}

// NewResourceConditionFuncInfo instantiates a ResourceConditionFuncInfo
func NewResourceConditionFuncInfo(node ast.Node, info *types.Info) *ResourceConditionFuncInfo {
	result := &ResourceConditionFuncInfo{
		TypesInfo: info,
	}

	switch node := node.(type) {
	case *ast.FuncDecl:
		result.AstFuncDecl = node
		result.Body = node.Body
		result.Node = node
		result.Pos = node.Pos()
		result.Type = node.Type
	case *ast.FuncLit:
		result.AstFuncLit = node
		result.Body = node.Body
		result.Node = node
		result.Pos = node.Pos()
		result.Type = node.Type
	}

	return result
}
