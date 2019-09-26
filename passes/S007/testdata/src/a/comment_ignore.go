package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S007
	_ = schema.Schema{
		Required:      true,
		ConflictsWith: []string{},
	}
}
