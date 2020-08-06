package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
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
