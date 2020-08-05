package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S007
	_ = schema.Schema{
		Required:      true,
		ConflictsWith: []string{},
	}
}
