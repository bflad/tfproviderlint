package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
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
