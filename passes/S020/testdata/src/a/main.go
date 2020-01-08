package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and enable ForceNew"
		Computed: true,
		ForceNew: true,
	}

	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		ForceNew: true,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and enable ForceNew"
			Computed: true,
			ForceNew: true,
		},
	}
}

func validateFunc(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
