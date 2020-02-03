package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S029
	_ = schema.Schema{
		Computed:     true,
		ExactlyOneOf: []string{"test"},
	}
}
