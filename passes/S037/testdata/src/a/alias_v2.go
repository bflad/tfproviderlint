package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
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
