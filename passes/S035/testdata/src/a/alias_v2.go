package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{
		AtLeastOneOf: []string{
			".invalidreference", // want "invalid AtLeastOneOf attribute reference"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			AtLeastOneOf: []string{
				".invalidreference", // want "invalid AtLeastOneOf attribute reference"
			},
		},
	}
}
