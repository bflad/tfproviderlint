package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S033
	_ = schema.Schema{
		Computed:  true,
		StateFunc: stateFunc_v2,
	}
}
