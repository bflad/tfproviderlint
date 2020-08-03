package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d s.ResourceData

	d.GetOkExists("test") // want "ResourceData.GetOkExists\\(\\) call should be avoided"

	d.Get("test")
	d.GetOk("test")
}
