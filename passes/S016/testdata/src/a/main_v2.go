package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema Set should only be included for TypeSet"
		Set:  set_v2,
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Set:  set_v2,
		Type: schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema Set should only be included for TypeSet"
			Set:  set_v2,
			Type: schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Set:  set_v2,
			Type: schema.TypeSet,
		},
	}
}

func set_v2(v interface{}) int {
	return 0
}
