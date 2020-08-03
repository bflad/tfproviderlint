package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S028
	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: defaultFunc_v2,
	}
}
