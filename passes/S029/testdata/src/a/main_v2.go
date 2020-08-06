package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure ExactlyOneOf"
		Computed:     true,
		ExactlyOneOf: []string{"test"},
	}

	_ = schema.Schema{
		Computed:     true,
		ExactlyOneOf: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		ExactlyOneOf: []string{"test"},
		Optional:     true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure ExactlyOneOf"
			Computed:     true,
			ExactlyOneOf: []string{"test"},
		},
	}
}
