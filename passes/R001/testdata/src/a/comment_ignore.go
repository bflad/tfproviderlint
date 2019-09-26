package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "") //lintignore:R001

	//lintignore:R001
	d.Set(key, "")
}
