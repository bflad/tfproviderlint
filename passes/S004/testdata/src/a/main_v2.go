package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not enable Required and configure Default"
		Required: true,
		Default:  true,
	}

	_ = schema.Schema{
		Required: true,
		Default:  nil,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = schema.Schema{
		Default: true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not enable Required and configure Default"
			Required: true,
			Default:  true,
		},
	}
}
