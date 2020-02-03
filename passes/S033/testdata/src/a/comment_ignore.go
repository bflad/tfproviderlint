package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S033
	_ = schema.Schema{
		Computed:  true,
		StateFunc: stateFunc,
	}
}
