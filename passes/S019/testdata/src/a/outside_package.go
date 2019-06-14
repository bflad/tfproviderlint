package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: false,
	}

	_ = schema.Schema{
		Optional: false,
	}

	_ = schema.Schema{
		Required: false,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: false,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Optional: false,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: false,
		},
	}
}
