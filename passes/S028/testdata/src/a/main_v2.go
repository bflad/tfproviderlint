package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure DefaultFunc"
		Computed:    true,
		DefaultFunc: defaultFunc_v2,
	}

	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		DefaultFunc: defaultFunc_v2,
		Optional:    true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure DefaultFunc"
			Computed:    true,
			DefaultFunc: defaultFunc_v2,
		},
	}
}

func defaultFunc_v2() (interface{}, error) {
	return nil, nil
}
