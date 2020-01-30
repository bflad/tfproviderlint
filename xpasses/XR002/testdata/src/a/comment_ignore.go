package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:XR002
	_ = schema.Resource{
		Create: createFunc,
	}
}
