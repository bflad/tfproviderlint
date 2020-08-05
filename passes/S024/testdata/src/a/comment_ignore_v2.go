package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S024
	_ = schema.Resource{
		Read: readFunc_v2,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}
