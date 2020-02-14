package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	// Comment ignored

	//lintignore:R008
	d.SetPartial("test")

	d.SetPartial("test") //lintignore:R008

	// Failing

	d.SetPartial("test") // want "deprecated \\(schema.ResourceData\\).SetPartial"

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	d.SetPartial("test") // want "deprecated \\(schema.ResourceData\\).SetPartial"

	return nil
}
