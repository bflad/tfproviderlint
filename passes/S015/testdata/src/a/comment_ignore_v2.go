package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S015
	_ = map[string]*schema.Schema{
		"INVALID": {},
	}

	//lintignore:S015
	_ = map[string]*schema.Schema{
		"invalid!": {},
	}

	//lintignore:S015
	_ = map[string]*schema.Schema{
		"invalid-name": {},
	}
}
