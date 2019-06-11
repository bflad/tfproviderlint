package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
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
