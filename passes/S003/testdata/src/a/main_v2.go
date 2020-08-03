package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not enable Required and Computed"
		Required: true,
		Computed: true,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not enable Required and Computed"
			Required: true,
			Computed: true,
		},
	}
}
