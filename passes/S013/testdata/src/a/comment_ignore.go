package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S013
	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
	}

	//lintignore:S013
	_ = map[string]*schema.Schema{
		"name": {
			Computed: false,
			Type:     schema.TypeString,
		},
	}

	//lintignore:S013
	_ = map[string]*schema.Schema{
		"name": {
			Optional: false,
			Type:     schema.TypeString,
		},
	}

	//lintignore:S013
	_ = map[string]*schema.Schema{
		"name": {
			Required: false,
			Type:     schema.TypeString,
		},
	}
}
