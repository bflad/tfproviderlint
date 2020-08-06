package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S010
	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: validateFunc_v2,
	}
}
