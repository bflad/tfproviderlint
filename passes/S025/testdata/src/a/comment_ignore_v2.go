package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S025
	_ = schema.Schema{
		AtLeastOneOf: []string{"test"},
		Computed:     true,
	}
}
