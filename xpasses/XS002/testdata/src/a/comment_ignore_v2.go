package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:XS002
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {},
			"arn":  {},
		},
	}
}
