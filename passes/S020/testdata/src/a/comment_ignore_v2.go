package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S020
	_ = schema.Schema{
		Computed: true,
		ForceNew: true,
	}
}
