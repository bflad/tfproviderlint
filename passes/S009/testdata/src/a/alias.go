package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         s.TypeList,
		ValidateFunc: validateFunc,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         s.TypeSet,
		ValidateFunc: validateFunc,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         s.TypeList,
			ValidateFunc: validateFunc,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         s.TypeSet,
			ValidateFunc: validateFunc,
		},
	}
}
