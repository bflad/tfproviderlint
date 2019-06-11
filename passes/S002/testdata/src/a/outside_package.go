package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required: true,
		Optional: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: true,
			Optional: true,
		},
	}
}
