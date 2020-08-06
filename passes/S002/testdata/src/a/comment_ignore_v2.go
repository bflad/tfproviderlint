package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S002
	_ = schema.Schema{
		Required: true,
		Optional: true,
	}
}
