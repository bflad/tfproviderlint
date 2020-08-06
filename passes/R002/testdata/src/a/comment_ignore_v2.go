package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) //lintignore:R002

	//lintignore:R002
	d.Set("test", *stringPtr)
}
