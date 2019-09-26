package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
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
