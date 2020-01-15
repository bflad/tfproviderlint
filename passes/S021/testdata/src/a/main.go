package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{
		Computed: true,
		ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
			ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed: true,
		},
	}
}
