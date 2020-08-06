package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure StateFunc"
		Computed:  true,
		StateFunc: stateFunc_v2,
	}

	_ = schema.Schema{
		Computed:  true,
		StateFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		StateFunc: stateFunc_v2,
		Optional:  true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure StateFunc"
			Computed:  true,
			StateFunc: stateFunc_v2,
		},
	}
}

func stateFunc_v2(interface{}) string {
	return ""
}
