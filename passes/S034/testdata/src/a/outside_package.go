package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		PromoteSingle: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			PromoteSingle: true,
		},
	}
}
