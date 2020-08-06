package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema of TypeMap should include Elem"
		Type: schema.TypeMap,
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Schema{Type: schema.TypeString},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: schema.TypeMap,
		},
	}
}
