package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S004
	_ = schema.Schema{
		Required: true,
		Default:  true,
	}
}
