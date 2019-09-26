package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{
		Computed: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

	_ = schema.Schema{
		Optional: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

	_ = schema.Schema{
		Required: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

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
			Computed: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Optional: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Required: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
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
