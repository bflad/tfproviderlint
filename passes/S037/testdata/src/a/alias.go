package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{
		ExactlyOneOf: []string{
			".invalidreference", // want "invalid ExactlyOneOf attribute reference"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			ExactlyOneOf: []string{
				".invalidreference", // want "invalid ExactlyOneOf attribute reference"
			},
		},
	}
}
