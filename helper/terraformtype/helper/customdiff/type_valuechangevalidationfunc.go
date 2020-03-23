package customdiff

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameValueChangeValidationFunc = `ValueChangeValidationFunc`
)

// IsFuncTypeValueChangeValidationFunc returns true if the FuncType matches expected parameters and results types
func IsFuncTypeValueChangeValidationFunc(node ast.Node, info *types.Info) bool {
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

	return astutils.IsFieldListType(funcType.Results, 0, astutils.IsExprTypeError)
}

// IsTypeValueChangeValidationFunc returns if the type is ValueChangeValidationFunc from the customdiff package
func IsTypeValueChangeValidationFunc(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameValueChangeValidationFunc)
	case *types.Pointer:
		return IsTypeValueChangeValidationFunc(t.Elem())
	default:
		return false
	}
}

// ValueChangeValidationFuncInfo represents all gathered ValueChangeValidationFunc data for easier access
type ValueChangeValidationFuncInfo struct {
	AstFuncDecl *ast.FuncDecl
	AstFuncLit  *ast.FuncLit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.FuncType
	TypesInfo   *types.Info
}

// NewValueChangeValidationFuncInfo instantiates a ValueChangeValidationFuncInfo
func NewValueChangeValidationFuncInfo(node ast.Node, info *types.Info) *ValueChangeValidationFuncInfo {
	result := &ValueChangeValidationFuncInfo{
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
