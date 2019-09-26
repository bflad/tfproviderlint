package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) //lintignore:R002

	//lintignore:R002
	d.Set("test", *stringPtr)
}
