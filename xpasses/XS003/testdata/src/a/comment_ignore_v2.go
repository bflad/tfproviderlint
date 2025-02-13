package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
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
