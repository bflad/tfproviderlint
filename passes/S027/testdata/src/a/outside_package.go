package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
		Default:  "test",
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			Default:  "test",
		},
	}
}
