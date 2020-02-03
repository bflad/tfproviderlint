package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure StateFunc"
		Computed:  true,
		StateFunc: stateFunc,
	}

	_ = schema.Schema{
		Computed:  true,
		StateFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		StateFunc: stateFunc,
		Optional:  true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure StateFunc"
			Computed:  true,
			StateFunc: stateFunc,
		},
	}
}

func stateFunc(interface{}) string {
	return ""
}
