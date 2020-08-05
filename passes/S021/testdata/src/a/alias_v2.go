package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{
		Computed:     true,
		ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
	}

	_ = s.Schema{
		Computed: true,
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed:     true,
			ComputedWhen: []string{"another_attr"}, // want "schema should omit ComputedWhen"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Computed: true,
		},
	}
}
