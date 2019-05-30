package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S004
	_ = schema.Schema{
		Required: true,
		Default:  true,
	}
}
