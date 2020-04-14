package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	validation.IsCIDRNetwork(0, 32)

	// Comment ignored

	//lintignore:V002
	validation.CIDRNetwork(0, 32)

	validation.CIDRNetwork(0, 32) //lintignore:V002

	// Failing

	validation.IsCIDRNetwork(0, 32) // want "deprecated validation.CIDRNetwork should be replaced with validation.IsCIDRNetwork"
}
