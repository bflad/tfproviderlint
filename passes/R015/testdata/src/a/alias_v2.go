package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d schema.ResourceData

	d.SetId(r.UniqueId()) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")
}
