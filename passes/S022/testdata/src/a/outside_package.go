package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Resource{},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
			Elem: &schema.Resource{},
		},
	}
}
