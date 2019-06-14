package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	//lintignore:S009
	_ = schema.Schema{
		Type:         schema.TypeList,
		ValidateFunc: validateFunc,
	}

	//lintignore:S009
	_ = schema.Schema{
		Type:         schema.TypeSet,
		ValidateFunc: validateFunc,
	}
}
