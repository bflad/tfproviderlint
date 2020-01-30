package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData

	d.GetOkExists("test") //lintignore:XR001

	//lintignore:XR001
	d.GetOkExists("test")
}
