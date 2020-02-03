package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
		MaxItems: 1,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			MaxItems: 1,
		},
	}
}
