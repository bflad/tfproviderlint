package a

import (
	t "time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	var d schema.ResourceData

	d.SetId(t.Now().Format(t.RFC3339)) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")
}
