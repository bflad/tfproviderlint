package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{
		ConflictsWith: []string{
			".invalidreference", // want "invalid ConflictsWith attribute reference"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			ConflictsWith: []string{
				".invalidreference", // want "invalid ConflictsWith attribute reference"
			},
		},
	}
}
