package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "") //lintignore:resourcedatasetkey

	//lintignore:resourcedatasetkey
	d.Set(key, "")
}
