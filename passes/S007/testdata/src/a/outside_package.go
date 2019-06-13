package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required:      true,
		ConflictsWith: []string{},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required:      true,
			ConflictsWith: []string{},
		},
	}
}
