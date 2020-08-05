package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S009
	_ = schema.Schema{
		Type:         schema.TypeList,
		ValidateFunc: validateFunc_v2,
	}

	//lintignore:S009
	_ = schema.Schema{
		Type:         schema.TypeSet,
		ValidateFunc: validateFunc_v2,
	}
}
