package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Default: true,
		Type:    schema.TypeList,
	}

	_ = schema.Schema{
		Default: true,
		Type:    schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Default: true,
			Type:    schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Default: true,
			Type:    schema.TypeSet,
		},
	}
}
