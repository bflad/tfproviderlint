package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required: true,
		Computed: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: true,
			Computed: true,
		},
	}
}
