package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
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
