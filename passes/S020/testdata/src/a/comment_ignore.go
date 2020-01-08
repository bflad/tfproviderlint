package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S020
	_ = schema.Schema{
		Computed: true,
		ForceNew: true,
	}
}
