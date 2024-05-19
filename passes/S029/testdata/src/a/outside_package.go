package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:     true,
		ExactlyOneOf: []string{"test"},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:     true,
			ExactlyOneOf: []string{"test"},
		},
	}
}
