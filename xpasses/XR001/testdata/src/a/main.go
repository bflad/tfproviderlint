package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")

	return nil
}
