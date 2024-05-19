package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: defaultFunc_v2,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:    true,
			DefaultFunc: defaultFunc_v2,
		},
	}
}
