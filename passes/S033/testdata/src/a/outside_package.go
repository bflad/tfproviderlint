package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:  true,
		StateFunc: stateFunc,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:  true,
			StateFunc: stateFunc,
		},
	}
}
