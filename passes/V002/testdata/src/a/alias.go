package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	v.IsCIDRNetwork(0, 32)

	// Comment ignored

	//lintignore:V002
	v.CIDRNetwork(0, 32)

	v.CIDRNetwork(0, 32) //lintignore:V002

	// Failing

	v.CIDRNetwork(0, 32) // want "deprecated v.CIDRNetwork should be replaced with v.IsCIDRNetwork"
}
