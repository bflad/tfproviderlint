package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		AtLeastOneOf: []string{"test"},
		Computed:     true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			AtLeastOneOf: []string{"test"},
			Computed:     true,
		},
	}
}
