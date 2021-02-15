package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema of TypeMap should include Elem"
		Type: schema.TypeMap,
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Schema{Type: schema.TypeString},
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: schemaFunc(),
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: schemaVar,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: schema.TypeMap,
		},
	}
}

func schemaFunc() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeString,
	}
}

var schemaVar = &schema.Schema{
	Type: schema.TypeString,
}
