package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema Set should only be included for TypeSet"
		Set:  set,
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Set:  set,
		Type: schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema Set should only be included for TypeSet"
			Set:  set,
			Type: schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Set:  set,
			Type: schema.TypeSet,
		},
	}
}

func set(v interface{}) int {
	return 0
}
