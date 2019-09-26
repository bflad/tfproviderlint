package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{
		Computed: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

	_ = s.Schema{
		Optional: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

	_ = s.Schema{
		Required: false, // want "schema should omit Computed, Optional, or Required set to false"
	}

	_ = s.Schema{
		Computed: true,
	}

	_ = s.Schema{
		Optional: true,
	}

	_ = s.Schema{
		Required: true,
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Optional: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Required: false, // want "schema should omit Computed, Optional, or Required set to false"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed: true,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Optional: true,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Required: true,
		},
	}
}
