package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S026
	_ = schema.Schema{
		Computed:      true,
		ConflictsWith: []string{"test"},
	}
}
