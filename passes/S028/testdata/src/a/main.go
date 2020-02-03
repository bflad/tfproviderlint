package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure DefaultFunc"
		Computed:    true,
		DefaultFunc: defaultFunc,
	}

	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		DefaultFunc: defaultFunc,
		Optional:    true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure DefaultFunc"
			Computed:    true,
			DefaultFunc: defaultFunc,
		},
	}
}

func defaultFunc() (interface{}, error) {
	return nil, nil
}
