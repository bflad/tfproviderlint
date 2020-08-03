package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// This first scenario is valid because Schema may be used within Elem
	// which should not require Description

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure Description"
			Optional: true,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Description: "test",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}
}
