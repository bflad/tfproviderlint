package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	d.GetOkExists("test") //lintignore:XR001

	//lintignore:XR001
	d.GetOkExists("test")
}
