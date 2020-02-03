package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S031
	_ = schema.Schema{
		Computed: true,
		MaxItems: 1,
	}
}
