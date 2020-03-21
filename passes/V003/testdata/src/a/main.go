package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.IsIPv4Range

	// Comment ignored

	//lintignore:V003
	validation.IPRange()

	validation.IPRange() //lintignore:V003

	// Failing

	_ = validation.IPRange // want "deprecated validation.IPRange should be replaced with validation.IsIPv4Range"
	validation.IPRange()   // want "deprecated validation.IPRange() should be replaced with validation.IsIPv4Range"
}
