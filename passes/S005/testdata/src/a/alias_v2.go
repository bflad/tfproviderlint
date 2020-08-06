package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema should not enable Computed and configure Default"
		Computed: true,
		Default:  true,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not enable Computed and configure Default"
			Computed: true,
			Default:  true,
		},
	}
}
