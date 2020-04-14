package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.IsIPAddress

	// Comment ignored

	//lintignore:V004
	v.SingleIP()

	v.SingleIP() //lintignore:V004

	// Failing

	_ = v.IsIPAddress // want "deprecated v.SingleIP should be replaced with v.IsIPAddress"
	v.IsIPAddress     // want "deprecated v.SingleIP() should be replaced with v.IsIPAddress"
}
