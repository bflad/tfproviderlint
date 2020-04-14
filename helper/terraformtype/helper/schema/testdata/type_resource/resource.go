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
					Elem: &schema.Resource{
					},
				},
			},
		}
	},
}
