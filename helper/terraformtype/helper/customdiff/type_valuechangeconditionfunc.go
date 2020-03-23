package customdiff

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameValueChangeConditionFunc = `ValueChangeConditionFunc`
)

// IsFuncTypeValueChangeConditionFunc returns true if the FuncType matches expected parameters and results types
func IsFuncTypeValueChangeConditionFunc(node ast.Node, info *types.Info) bool {
	funcType := astutils.FuncTypeFromNode(node)

	if funcType == nil {
		return false
	}

	if !astutils.HasFieldListLength(funcType.Params, 3) {
		return false
	}

	if !astutils.IsFieldListType(funcType.Params, 0, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.IsFieldListType(funcType.Params, 1, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.IsFieldListType(funcType.Params, 2, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.HasFieldListLength(funcType.Results, 1) {
		return false
	}

	return astutils.IsFieldListType(funcType.Results, 0, astutils.IsExprTypeBool)
}

// IsTypeValueChangeConditionFunc returns if the type is ValueChangeConditionFunc from the customdiff package
func IsTypeValueChangeConditionFunc(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameValueChangeConditionFunc)
	case *types.Pointer:
		return IsTypeValueChangeConditionFunc(t.Elem())
	default:
		return false
	}
}

// ValueChangeConditionFuncInfo represents all gathered ValueChangeConditionFunc data for easier access
type ValueChangeConditionFuncInfo struct {
	AstFuncDecl *ast.FuncDecl
	AstFuncLit  *ast.FuncLit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.FuncType
	TypesInfo   *types.Info
}

// NewValueChangeConditionFuncInfo instantiates a ValueChangeConditionFuncInfo
func NewValueChangeConditionFuncInfo(node ast.Node, info *types.Info) *ValueChangeConditionFuncInfo {
	result := &ValueChangeConditionFuncInfo{
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
