package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S016
	_ = schema.Schema{
		Set:  set,
		Type: schema.TypeList,
	}

	//lintignore:S016
	_ = map[string]*schema.Schema{
		"name": {
			Set:  set,
			Type: schema.TypeList,
		},
	}
}
