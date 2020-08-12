package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	d.SetId(time.Now().Format(time.RFC3339)) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	d.SetId(time.Now().Format(time.RFC3339)) // want "schema attributes should be stable across Terraform runs"

	d.SetId("test")

	return nil
}
