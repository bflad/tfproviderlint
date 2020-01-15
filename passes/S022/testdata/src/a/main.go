package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Resource{}, // want "schema of TypeMap should not use Elem of \\*schema.Resource"
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Schema{},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
			Elem: &schema.Resource{}, // want "schema of TypeMap should not use Elem of \\*schema.Resource"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
			Elem: &schema.Schema{},
		},
	}
}
