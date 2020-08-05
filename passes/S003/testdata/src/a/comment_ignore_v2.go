package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S003
	_ = schema.Schema{
		Required: true,
		Computed: true,
	}
}
