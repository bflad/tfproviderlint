package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:XS001
	_ = map[string]*schema.Schema{
		"name": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
