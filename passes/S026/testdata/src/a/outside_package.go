package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:      true,
		ConflictsWith: []string{"test"},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:      true,
			ConflictsWith: []string{"test"},
		},
	}
}
