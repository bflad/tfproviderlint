package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure DiffSuppressFunc"
		Computed:         true,
		DiffSuppressFunc: diffSuppressFunc,
	}

	_ = schema.Schema{
		Computed:         true,
		DiffSuppressFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		DiffSuppressFunc: diffSuppressFunc,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure DiffSuppressFunc"
			Computed:         true,
			DiffSuppressFunc: diffSuppressFunc,
		},
	}
}

func diffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	return true
}
