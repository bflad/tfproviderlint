package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S023
	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeBool,
	}

	//lintignore:S023
	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeFloat,
	}

	//lintignore:S023
	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeInt,
	}

	//lintignore:S023
	_ = schema.Schema{
		Elem: &schema.Schema{},
		Type: schema.TypeString,
	}

	//lintignore:S023
	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{},
			Type: schema.TypeBool,
		},
	}

	//lintignore:S023
	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{},
			Type: schema.TypeFloat,
		},
	}

	//lintignore:S023
	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{},
			Type: schema.TypeInt,
		},
	}

	//lintignore:S023
	_ = map[string]*schema.Schema{
		"name": {
			Elem: &schema.Schema{},
			Type: schema.TypeString,
		},
	}
}
