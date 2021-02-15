package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: schema.TypeList,
	}

	_ = schema.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: schema.TypeSet,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Schema{Type: schema.TypeString},
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: &schema.Schema{Type: schema.TypeString},
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: resourceSchemaFunc(),
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: resourceSchemaFunc(),
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: resourceSchemaVar,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: resourceSchemaVar,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: schemaVar,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: schemaVar,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: schemaFunc(),
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: schemaFunc(),
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should include Elem"
			Type: schema.TypeList,
		},
	}
}

func resourceSchemaFunc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"test": {
				Type: schema.TypeString,
			},
		},
	}
}

var resourceSchemaVar = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"test": {
			Type: schema.TypeString,
		},
	},
}

func schemaFunc() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeString,
	}
}

var schemaVar = &schema.Schema{
	Type: schema.TypeString,
}
