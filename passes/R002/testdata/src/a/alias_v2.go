package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d s.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)
}
