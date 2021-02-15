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

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: schemaFunc_v2(),
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: schemaVar_v2,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: schema.TypeMap,
		},
	}
}

func schemaFunc_v2() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeString,
	}
}

var schemaVar_v2 = &schema.Schema{
	Type: schema.TypeString,
}
