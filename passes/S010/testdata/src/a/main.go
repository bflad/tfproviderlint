package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure ValidateFunc"
		Computed:     true,
		ValidateFunc: validateFunc,
	}

	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		ValidateFunc: validateFunc,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure ValidateFunc"
			Computed:     true,
			ValidateFunc: validateFunc,
		},
	}
}

func validateFunc(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
