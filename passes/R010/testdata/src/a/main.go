package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData
	var o, n interface{}

	// Comment ignored

	//lintignore:R010
	_, n = d.GetChange("test")

	_, n = d.GetChange("test") //lintignore:R010

	// Failing

	_, n = d.GetChange("test") // want "prefer d.Get\\(\\) over d.GetChange\\(\\) when only using second return value"

	// Passing

	o, n = d.GetChange("test")

	o, _ = d.GetChange("test")

	// Prevent these go build errors:
	// o declared but not used
	// n declared but not used
	_ = o
	_ = n
}
