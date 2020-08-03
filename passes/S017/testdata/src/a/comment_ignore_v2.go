package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S017
	_ = schema.Schema{
		MaxItems: 1,
		Type:     schema.TypeString,
	}

	//lintignore:S017
	_ = schema.Schema{
		MinItems: 1,
		Type:     schema.TypeString,
	}

	//lintignore:S017
	_ = map[string]*schema.Schema{
		"name": {
			MaxItems: 1,
			Type:     schema.TypeString,
		},
	}

	//lintignore:S017
	_ = map[string]*schema.Schema{
		"name": {
			MinItems: 1,
			Type:     schema.TypeString,
		},
	}
}
