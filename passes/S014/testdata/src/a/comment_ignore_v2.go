package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S014
	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			},
			Computed: true,
			Type:     schema.TypeList,
		},
	}
	//lintignore:S014
	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			Optional: true,
			Type:     schema.TypeList,
		},
	}

	//lintignore:S014
	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			Required: true,
			Type:     schema.TypeList,
		},
	}
}
