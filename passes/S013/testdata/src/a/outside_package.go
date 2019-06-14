package a

import (
	"a/schema"
)

func foutside() {
	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: false,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Optional: false,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: false,
			Type:     schema.TypeString,
		},
	}
}
