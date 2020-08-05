package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S022
	_ = schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Resource{},
	}

	//lintignore:S022
	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
			Elem: &schema.Resource{},
		},
	}
}
