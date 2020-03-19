package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		AtLeastOneOf: []string{
			".invalidreference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			AtLeastOneOf: []string{
				".invalidreference",
			},
		},
	}
}
