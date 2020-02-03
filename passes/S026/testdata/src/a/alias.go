package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not only enable Computed and configure ConflictsWith"
		Computed:      true,
		ConflictsWith: []string{"test"},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not only enable Computed and configure ConflictsWith"
			Computed:      true,
			ConflictsWith: []string{"test"},
		},
	}
}
