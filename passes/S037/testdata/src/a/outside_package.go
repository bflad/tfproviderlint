package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		ExactlyOneOf: []string{
			".invalidreference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ExactlyOneOf: []string{
				".invalidreference",
			},
		},
	}
}
