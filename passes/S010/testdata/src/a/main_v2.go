package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure ValidateFunc"
		Computed:     true,
		ValidateFunc: validateFunc_v2,
	}

	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		ValidateFunc: validateFunc_v2,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure ValidateFunc"
			Computed:     true,
			ValidateFunc: validateFunc_v2,
		},
	}
}

func validateFunc_v2(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
