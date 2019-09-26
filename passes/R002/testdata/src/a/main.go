package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)
	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)

	return nil
}
