package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
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
