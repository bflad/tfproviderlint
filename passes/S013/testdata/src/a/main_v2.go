package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// These first scenarios are valid because Schema may be used within Elem
	// which does not require one of Computed, Optional, or Required

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = schema.Schema{
		Computed: false,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Optional: false,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Required: false,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Computed: true,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Optional: true,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Required: true,
		Type:     schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Type: schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Computed: false,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Optional: false,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure one of Computed, Optional, or Required"
			Required: false,
			Type:     schema.TypeString,
		},
	}
}
