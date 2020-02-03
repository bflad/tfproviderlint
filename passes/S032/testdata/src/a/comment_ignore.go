package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S032
	_ = schema.Schema{
		Computed: true,
		MinItems: 1,
	}
}
