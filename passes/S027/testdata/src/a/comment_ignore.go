package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S027
	_ = schema.Schema{
		Computed: true,
		Default:  "test",
	}
}
