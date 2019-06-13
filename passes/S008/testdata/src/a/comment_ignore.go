package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S008
	_ = schema.Schema{
		Default: true,
		Type:    schema.TypeList,
	}

	//lintignore:S008
	_ = schema.Schema{
		Default: true,
		Type:    schema.TypeSet,
	}
}
