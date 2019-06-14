package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		Optional: true,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Optional: true,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: true,
		},
	}
}
