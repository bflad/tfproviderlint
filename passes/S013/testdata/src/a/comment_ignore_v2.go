package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
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
