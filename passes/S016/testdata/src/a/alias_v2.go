package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema Set should only be included for TypeSet"
		Set:  set_v2,
		Type: s.TypeList,
	}

	_ = s.Schema{
		Set:  set_v2,
		Type: s.TypeSet,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema Set should only be included for TypeSet"
			Set:  set_v2,
			Type: s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Set:  set_v2,
			Type: s.TypeSet,
		},
	}
}
