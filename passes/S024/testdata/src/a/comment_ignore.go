package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S024
	_ = schema.Resource{
		Read: readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}
