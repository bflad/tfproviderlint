package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:    schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}
}
