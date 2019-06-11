package terraformtype

import (
	"go/types"
	"strings"
)

func IsTypeHelperSchema(t types.Type) bool {
	switch t := t.(type) {
	default:
		return false
	case *types.Named:
		if t.Obj().Name() != "Schema" {
			return false
		}
		// HasSuffix here due to vendoring
		if !strings.HasSuffix(t.Obj().Pkg().Path(), "github.com/hashicorp/terraform/helper/schema") {
			return false
		}
	}
	return true
}
