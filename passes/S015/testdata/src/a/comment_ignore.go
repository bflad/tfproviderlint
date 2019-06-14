package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
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
