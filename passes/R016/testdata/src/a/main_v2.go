package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData

	d.SetId(resource.PrefixedUniqueId("test")) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	d.SetId(resource.PrefixedUniqueId("test")) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	return nil
}
