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

	d.Partial(true)  // want "deprecated \\(schema.ResourceData\\).Partial"
	d.Partial(false) // want "deprecated \\(schema.ResourceData\\).Partial"

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)  // want "deprecated \\(schema.ResourceData\\).Partial"
	d.Partial(false) // want "deprecated \\(schema.ResourceData\\).Partial"

	return nil
}
