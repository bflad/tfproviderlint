package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S003
	_ = schema.Schema{
		Required: true,
		Computed: true,
	}
}
