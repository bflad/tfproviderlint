package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	// Comment ignored

	//lintignore:R007
	d.Partial(true)

	d.Partial(true) //lintignore:R007

	// Failing

	// want "deprecated d.Partial"
	// want "deprecated d.Partial"

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	// want "deprecated d.Partial"
	// want "deprecated d.Partial"

	return nil
}
