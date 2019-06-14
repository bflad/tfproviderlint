package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeSet,
		},
	}
}
