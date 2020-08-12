package a

import (
	t "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d schema.ResourceData

	d.SetId(t.Now().Format(t.RFC3339)) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")
}
