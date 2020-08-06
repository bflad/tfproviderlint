package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:XS001
	_ = map[string]*schema.Schema{
		"name": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
