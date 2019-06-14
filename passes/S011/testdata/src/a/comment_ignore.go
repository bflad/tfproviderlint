package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S011
	_ = schema.Schema{
		Computed:         true,
		DiffSuppressFunc: diffSuppressFunc,
	}
}
