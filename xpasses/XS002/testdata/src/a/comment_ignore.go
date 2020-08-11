package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:XS002
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {},
			"arn":  {},
		},
	}
}
