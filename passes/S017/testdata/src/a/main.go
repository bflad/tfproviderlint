package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
		MaxItems: 1,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{ // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
		MinItems: 1,
		Type:     schema.TypeString,
	}

	_ = schema.Schema{
		Type: schema.TypeString,
	}

	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeList,
	}

	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeMap,
	}

	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeSet,
	}

	_ = schema.Schema{
		MinItems: 1,
		Type:     schema.TypeList,
	}

	_ = schema.Schema{
		MinItems: 1,
		Type:     schema.TypeMap,
	}

	_ = schema.Schema{
		MinItems: 1,
		Type:     schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
			MaxItems: 1,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
			MinItems: 1,
			Type:     schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeMap,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeSet,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MinItems: 1,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MinItems: 1,
			Type:     schema.TypeMap,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MinItems: 1,
			Type:     schema.TypeSet,
		},
	}
}
