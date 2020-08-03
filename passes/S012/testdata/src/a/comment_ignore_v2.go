package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S012
	_ = schema.Schema{
		Computed: true,
	}

	//lintignore:S012
	_ = schema.Schema{
		Optional: true,
	}

	//lintignore:S012
	_ = schema.Schema{
		Required: true,
	}
}
