package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S011
	_ = schema.Schema{
		Computed:         true,
		DiffSuppressFunc: diffSuppressFunc,
	}
}
