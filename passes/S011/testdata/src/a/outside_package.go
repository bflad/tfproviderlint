package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:         true,
		DiffSuppressFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:         true,
			DiffSuppressFunc: func() {},
		},
	}
}
