package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
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
