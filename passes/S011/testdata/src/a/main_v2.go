package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema should not only enable Computed and configure DiffSuppressFunc"
		Computed:         true,
		DiffSuppressFunc: diffSuppressFunc_v2,
	}

	_ = schema.Schema{
		Computed:         true,
		DiffSuppressFunc: nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		DiffSuppressFunc: diffSuppressFunc_v2,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not only enable Computed and configure DiffSuppressFunc"
			Computed:         true,
			DiffSuppressFunc: diffSuppressFunc_v2,
		},
	}
}

func diffSuppressFunc_v2(k, old, new string, d *schema.ResourceData) bool {
	return true
}
