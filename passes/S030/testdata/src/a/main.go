package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure InputDefault"
		Computed:     true,
		InputDefault: "test",
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		InputDefault: "test",
		Optional:     true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure InputDefault"
			Computed:     true,
			InputDefault: "test",
		},
	}
}
