package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure AtLeastOneOf"
		AtLeastOneOf: []string{"test"},
		Computed:     true,
	}

	_ = schema.Schema{
		AtLeastOneOf: nil,
		Computed:     true,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		AtLeastOneOf: []string{"test"},
		Optional:     true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure AtLeastOneOf"
			AtLeastOneOf: []string{"test"},
			Computed:     true,
		},
	}
}
