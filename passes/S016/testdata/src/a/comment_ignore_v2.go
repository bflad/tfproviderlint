package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S016
	_ = schema.Schema{
		Set:  set_v2,
		Type: schema.TypeList,
	}

	//lintignore:S016
	_ = map[string]*schema.Schema{
		"name": {
			Set:  set_v2,
			Type: schema.TypeList,
		},
	}
}
