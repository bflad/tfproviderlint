package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not only enable Computed and configure ValidateFunc"
		Computed:     true,
		ValidateFunc: validateFunc,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not only enable Computed and configure ValidateFunc"
			Computed:     true,
			ValidateFunc: validateFunc,
		},
	}
}
