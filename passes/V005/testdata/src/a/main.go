package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.StringIsJSON

	// Comment ignored

	//lintignore:V005
	_ = validation.ValidateJsonString

	_ = validation.ValidateJsonString //lintignore:V005

	// Failing

	_ = validation.ValidateJsonString // want "deprecated validation.ValidateJsonString should be replaced with validation.StringIsJSON"
}
