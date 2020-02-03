package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
		MinItems: 1,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			MinItems: 1,
		},
	}
}
