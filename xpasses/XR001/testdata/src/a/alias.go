package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	var d s.ResourceData

	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")
}
