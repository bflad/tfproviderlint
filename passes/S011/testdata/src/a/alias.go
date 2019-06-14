package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not only enable Computed and configure DiffSuppressFunc"
		Computed:         true,
		DiffSuppressFunc: diffSuppressFunc,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not only enable Computed and configure DiffSuppressFunc"
			Computed:         true,
			DiffSuppressFunc: diffSuppressFunc,
		},
	}
}
