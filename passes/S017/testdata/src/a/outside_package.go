package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		MinItems: 1,
		Type:     schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MinItems: 1,
			Type:     schema.TypeString,
		},
	}
}
