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

	// want "deprecated d.SetPartial"

	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	// want "deprecated d.SetPartial"

	return nil
}
