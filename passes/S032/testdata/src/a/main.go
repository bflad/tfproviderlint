package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure MinItems"
		Computed: true,
		MinItems: 1,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		MinItems: 1,
		Optional: true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure MinItems"
			Computed: true,
			MinItems: 1,
		},
	}
}
