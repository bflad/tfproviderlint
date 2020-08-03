package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	// These first scenarios are valid because Schema may be used within Elem
	// which does not require one of Computed, Optional, or Required

	_ = s.Schema{
		Type: s.TypeString,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Type: s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Computed: false,
			Type:     s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Optional: false,
			Type:     s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Required: false,
			Type:     s.TypeString,
		},
	}
}
