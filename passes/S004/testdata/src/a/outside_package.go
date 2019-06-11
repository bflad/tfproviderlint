package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required: true,
		Default:  true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: true,
			Default:  true,
		},
	}
}
