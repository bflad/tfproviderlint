package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	d.SetId(resource.UniqueId()) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	d.SetId(resource.UniqueId()) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	return nil
}
