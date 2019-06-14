package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should configure Type"
		Computed: true,
	}

	_ = s.Schema{ // want "schema should configure Type"
		Optional: true,
	}

	_ = s.Schema{ // want "schema should configure Type"
		Required: true,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure Type"
			Computed: true,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure Type"
			Optional: true,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should configure Type"
			Required: true,
		},
	}
}
