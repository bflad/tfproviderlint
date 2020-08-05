package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)
	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	var stringPtr *string

	d.Set("test", *stringPtr) // want "ResourceData.Set\\(\\) pointer value dereference is extraneous"

	d.Set("test", stringPtr)

	return nil
}
