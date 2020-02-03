package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S028
	_ = schema.Schema{
		Computed:    true,
		DefaultFunc: defaultFunc,
	}
}
