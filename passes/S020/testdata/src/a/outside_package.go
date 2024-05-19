package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
		ForceNew: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			ForceNew: true,
		},
	}
}
