package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S018
	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeSet,
	}

	//lintignore:S018
	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeSet,
		},
	}
}
