package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema Set should only be included for TypeSet"
		Set:  set,
		Type: s.TypeList,
	}

	_ = s.Schema{
		Set:  set,
		Type: s.TypeSet,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema Set should only be included for TypeSet"
			Set:  set,
			Type: s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Set:  set,
			Type: s.TypeSet,
		},
	}
}
