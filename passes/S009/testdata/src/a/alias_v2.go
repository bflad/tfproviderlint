package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         s.TypeList,
		ValidateFunc: validateFunc_v2,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         s.TypeSet,
		ValidateFunc: validateFunc_v2,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         s.TypeList,
			ValidateFunc: validateFunc_v2,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         s.TypeSet,
			ValidateFunc: validateFunc_v2,
		},
	}
}
