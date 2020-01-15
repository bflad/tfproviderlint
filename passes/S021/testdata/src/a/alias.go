package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{
		Computed: true,
		ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
	}

	_ = s.Schema{
		Computed: true,
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed: true,
			ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed: true,
		},
	}
}
