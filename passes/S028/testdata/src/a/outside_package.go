package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: defaultFunc,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:    true,
			DefaultFunc: defaultFunc,
		},
	}
}
