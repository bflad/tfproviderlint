package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S010
	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: validateFunc,
	}
}
