package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		ConflictsWith: []string{
			".invalidreference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ConflictsWith: []string{
				".invalidreference",
			},
		},
	}
}
