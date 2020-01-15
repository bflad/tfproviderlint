package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeBool,
	}

	_ = schema.Schema{
		Type: schema.TypeBool,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeFloat,
	}

	_ = schema.Schema{
		Type: schema.TypeFloat,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeInt,
	}

	_ = schema.Schema{
		Type: schema.TypeInt,
	}

	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeString,
	}

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{},
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
			Elem: &schema.Schema{},
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
			Elem: &schema.Schema{},
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
			Elem: &schema.Schema{},
			Type: schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
	}
}
