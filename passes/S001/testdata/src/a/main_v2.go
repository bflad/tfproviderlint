package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
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
		Elem: resourceSchemaFunc_v2(),
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: resourceSchemaFunc_v2(),
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: resourceSchemaVar_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: resourceSchemaVar_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: schemaVar_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: schemaVar_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: schemaFunc_v2(),
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: schemaFunc_v2(),
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should include Elem"
			Type: schema.TypeList,
		},
	}
}

func resourceSchemaFunc_v2() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"test": {
				Type: schema.TypeString,
			},
		},
	}
}

var resourceSchemaVar_v2 = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"test": {
			Type: schema.TypeString,
		},
	},
}

func schemaFunc_v2() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeString,
	}
}

var schemaVar_v2 = &schema.Schema{
	Type: schema.TypeString,
}
