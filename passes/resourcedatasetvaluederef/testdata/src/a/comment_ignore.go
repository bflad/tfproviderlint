package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) //lintignore:resourcedatasetvaluederef

	//lintignore:resourcedatasetvaluederef
	d.Set("test", *stringPtr)
}
