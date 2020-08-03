package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not enable Required and configure ConflictsWith"
		Required:      true,
		ConflictsWith: []string{},
	}

	_ = schema.Schema{
		Required:      true,
		ConflictsWith: nil,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = schema.Schema{
		ConflictsWith: []string{},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not enable Required and configure ConflictsWith"
			Required:      true,
			ConflictsWith: []string{},
		},
	}
}
