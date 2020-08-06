package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S031
	_ = schema.Schema{
		Computed: true,
		MaxItems: 1,
	}
}
