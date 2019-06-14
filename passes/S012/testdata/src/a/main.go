package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should configure Type"
		Computed: true,
	}

	_ = schema.Schema{ // want "schema should configure Type"
		Optional: true,
	}

	_ = schema.Schema{ // want "schema should configure Type"
		Required: true,
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
		"name": { // want "schema should configure Type"
			Computed: true,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure Type"
			Optional: true,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should configure Type"
			Required: true,
		},
	}
}
