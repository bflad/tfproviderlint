package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S002
	_ = schema.Schema{
		Required: true,
		Optional: true,
	}
}
