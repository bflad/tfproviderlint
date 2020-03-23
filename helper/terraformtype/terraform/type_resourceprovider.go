package terraform

import (
	"go/types"
)

const (
	TypeNameResourceProvider = `ResourceProvider`
)

// IsTypeResourceProvider returns if the type is ResourceProvider from the terraform package
func IsTypeResourceProvider(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameResourceProvider)
	case *types.Pointer:
		return IsTypeResourceProvider(t.Elem())
	default:
		return false
	}
}
