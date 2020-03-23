package customdiff

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameValueValidationFunc = `ValueValidationFunc`
)

// IsFuncTypeValueValidationFunc returns true if the FuncType matches expected parameters and results types
func IsFuncTypeValueValidationFunc(node ast.Node, info *types.Info) bool {
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

	return astutils.IsFieldListType(funcType.Results, 0, astutils.IsExprTypeError)
}

// IsTypeValueValidationFunc returns if the type is ValueValidationFunc from the customdiff package
func IsTypeValueValidationFunc(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameValueValidationFunc)
	case *types.Pointer:
		return IsTypeValueValidationFunc(t.Elem())
	default:
		return false
	}
}

// ValueValidationFuncInfo represents all gathered ValueValidationFunc data for easier access
type ValueValidationFuncInfo struct {
	AstFuncDecl *ast.FuncDecl
	AstFuncLit  *ast.FuncLit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.FuncType
	TypesInfo   *types.Info
}

// NewValueValidationFuncInfo instantiates a ValueValidationFuncInfo
func NewValueValidationFuncInfo(node ast.Node, info *types.Info) *ValueValidationFuncInfo {
	result := &ValueValidationFuncInfo{
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
