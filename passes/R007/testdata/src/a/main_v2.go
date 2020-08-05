package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData

	// Comment ignored

	//lintignore:R007
	d.Partial(true)

	d.Partial(true) //lintignore:R007

	// Failing

	d.Partial(true)  // want "deprecated d.Partial"
	d.Partial(false) // want "deprecated d.Partial"

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)  // want "deprecated d.Partial"
	d.Partial(false) // want "deprecated d.Partial"

	return nil
}
