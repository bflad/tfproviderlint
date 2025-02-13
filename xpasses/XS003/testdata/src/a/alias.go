package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{
		Schema: map[string]*s.Schema{
			"a": { // want "XS003: schema might introduce crash or diff as it allows to be an empty block"
				Optional: true,
				Type:     s.TypeList,
				Elem: &s.Resource{
					Schema: map[string]*s.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}
}
