package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type: schema.TypeMap,
	}
}
