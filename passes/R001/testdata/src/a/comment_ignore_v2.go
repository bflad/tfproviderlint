package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "") //lintignore:R001

	//lintignore:R001
	d.Set(key, "")
}
