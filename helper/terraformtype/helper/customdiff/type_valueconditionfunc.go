package customdiff

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameValueConditionFunc = `ValueConditionFunc`
)

// IsFuncTypeValueConditionFunc returns true if the FuncType matches expected parameters and results types
func IsFuncTypeValueConditionFunc(node ast.Node, info *types.Info) bool {
	funcType := astutils.FuncTypeFromNode(node)

	if funcType == nil {
		return false
	}

	if !astutils.HasFieldListLength(funcType.Params, 2) {
		return false
	}

	if !astutils.IsFieldListType(funcType.Params, 0, astutils.IsExprTypeInterface) {
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

// IsTypeValueConditionFunc returns if the type is ValueConditionFunc from the customdiff package
func IsTypeValueConditionFunc(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameValueConditionFunc)
	case *types.Pointer:
		return IsTypeValueConditionFunc(t.Elem())
	default:
		return false
	}
}

// ValueConditionFuncInfo represents all gathered ValueConditionFunc data for easier access
type ValueConditionFuncInfo struct {
	AstFuncDecl *ast.FuncDecl
	AstFuncLit  *ast.FuncLit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.FuncType
	TypesInfo   *types.Info
}

// NewValueConditionFuncInfo instantiates a ValueConditionFuncInfo
func NewValueConditionFuncInfo(node ast.Node, info *types.Info) *ValueConditionFuncInfo {
	result := &ValueConditionFuncInfo{
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
