package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeSet,
		},
	}
}
