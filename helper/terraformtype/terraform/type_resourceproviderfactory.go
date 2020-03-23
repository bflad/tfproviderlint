package terraform

import (
	"go/types"
)

const (
	TypeNameResourceProviderFactory = `ResourceProviderFactory`
)

// IsTypeResourceProviderFactory returns if the type is ResourceProviderFactory from the terraform package
func IsTypeResourceProviderFactory(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameResourceProviderFactory)
	case *types.Pointer:
		return IsTypeResourceProviderFactory(t.Elem())
	default:
		return false
	}
}
