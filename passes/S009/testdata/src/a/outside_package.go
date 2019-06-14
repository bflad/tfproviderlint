package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type:         schema.TypeList,
		ValidateFunc: func() {},
	}

	_ = schema.Schema{
		Type:         schema.TypeSet,
		ValidateFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeList,
			ValidateFunc: func() {},
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeSet,
			ValidateFunc: func() {},
		},
	}
}
