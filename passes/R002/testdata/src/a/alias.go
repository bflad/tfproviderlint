package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	var d s.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)
}
