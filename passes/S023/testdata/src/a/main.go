package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{
		Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: schema.TypeBool,
	}

	_ = schema.Schema{
		Type: schema.TypeBool,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: schema.TypeFloat,
	}

	_ = schema.Schema{
		Type: schema.TypeFloat,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: schema.TypeInt,
	}

	_ = schema.Schema{
		Type: schema.TypeInt,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: schema.TypeString,
	}

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: schema.TypeBool,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeBool,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: schema.TypeFloat,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeFloat,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: schema.TypeInt,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeInt,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
	}
}
