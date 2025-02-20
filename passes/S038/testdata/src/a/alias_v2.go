package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{
		Type:    s.TypeString,
		Default: true, // want "schema should not declare Default with incompatible Type"
	}

	_ = map[string]*s.Schema{
		"name": {
			Type:    s.TypeString,
			Default: true, // want "schema should not declare Default with incompatible Type"
		},
	}
}
