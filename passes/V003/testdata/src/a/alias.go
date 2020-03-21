package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.IsIPv4Range

	// Comment ignored

	//lintignore:V003
	v.IPRange()

	v.IPRange() //lintignore:V003

	// Failing

	_ = v.IPRange // want "deprecated v.IPRange should be replaced with v.IsIPv4Range"
	v.IPRange()   // want "deprecated v.IPRange() should be replaced with v.IsIPv4Range"
}
