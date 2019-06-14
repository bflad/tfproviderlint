package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should use TypeList with MaxItems 1"
		MaxItems: 1,
		Type:     schema.TypeSet,
	}

	_ = schema.Schema{
		MaxItems: 2,
		Type:     schema.TypeSet,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
	}

	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeList,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should use TypeList with MaxItems 1"
			MaxItems: 1,
			Type:     schema.TypeSet,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 2,
			Type:     schema.TypeSet,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeSet,
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeList,
		},
	}
}
