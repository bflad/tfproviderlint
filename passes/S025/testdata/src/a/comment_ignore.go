package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S025
	_ = schema.Schema{
		AtLeastOneOf: []string{"test"},
		Computed:     true,
	}
}
