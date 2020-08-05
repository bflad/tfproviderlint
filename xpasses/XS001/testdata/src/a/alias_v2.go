package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	// This first scenario is valid because Schema may be used within Elem
	// which should not require Description

	_ = s.Schema{
		Optional: true,
		Type:     s.TypeString,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure Description"
			Optional: true,
			Type:     s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Description: "test",
			Optional:    true,
			Type:        s.TypeString,
		},
	}
}
