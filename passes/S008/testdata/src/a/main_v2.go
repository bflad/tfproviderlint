package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include Default"
		Default: true,
		Type:    schema.TypeList,
	}

	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include Default"
		Default: true,
		Type:    schema.TypeSet,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include Default"
			Default: true,
			Type:    schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include Default"
			Default: true,
			Type:    schema.TypeSet,
		},
	}
}
