package type_schema

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type schemaFunc func() *schema.Schema

var SchemaFuncs = []schemaFunc{

	// empty schema
	func() *schema.Schema {
		return &schema.Schema{}
	},

	// simple schema
	func() *schema.Schema {
		return &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		}
	},

	// empty nested block
	func() *schema.Schema {
		return &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     &schema.Resource{},
		}
	},

	// nested block
	func() *schema.Schema {
		return &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"foo": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  0,
					},
				},
			},
		}
	},
}
