package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
		Default:  true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			Default:  true,
		},
	}
}
