package type_resource

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type resourceFunc func() *schema.Resource

var ResourceFuncs = []resourceFunc{

	// empty resource
	func() *schema.Resource {
		return &schema.Resource{}
	},

	// resource with nested schema
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"foo": {},
			},
		}
	},

	// resource with nested schema which has nested resource
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"foo": {
					Type:     schema.TypeList,
					Required: true,
					MaxItems: 1,
					Elem:     &schema.Resource{},
				},
			},
		}
	},

	// resource with nested schema which has nested schema
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"test": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		}
	},

	// Ensure no panic with constant keys
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				ConstSchemaKeyTest: {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		}
	},

	// Ensure no panic with variable keys
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				VarSchemaKeyTest: {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		}
	},

	// Ensure no panic with function call keys
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				FuncSchemaKey(): {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		}
	},

	// Ensure no panic with variable values
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"test": VarSchemaValue,
			},
		}
	},

	// Ensure no panic with function call values
	func() *schema.Resource {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"test": FuncSchemaValue(),
			},
		}
	},
}

const (
	ConstSchemaKeyTest = "test"
)

var (
	VarSchemaKeyTest = "test"
	VarSchemaValue   = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
)

func FuncSchemaKey() string {
	return "test"
}

func FuncSchemaValue() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
}
