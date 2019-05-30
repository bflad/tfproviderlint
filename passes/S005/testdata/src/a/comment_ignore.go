package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S005
	_ = schema.Schema{
		Computed: true,
		Default:  true,
	}
}
