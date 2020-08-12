package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	var d schema.ResourceData

	d.SetId(r.PrefixedUniqueId("test")) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")
}
