package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	//lintignore:XR003
	_ = schema.Resource{
		Create: createFunc,
	}
}
