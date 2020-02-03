package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:S030
	_ = schema.Schema{
		Computed:     true,
		InputDefault: "test",
	}
}
