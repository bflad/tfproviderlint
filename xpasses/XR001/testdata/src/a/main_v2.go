package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData

	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")

	return nil
}
