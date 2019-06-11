package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type: schema.TypeMap,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
		},
	}
}
