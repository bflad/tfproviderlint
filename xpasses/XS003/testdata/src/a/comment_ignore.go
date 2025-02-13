package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:XS003
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
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
