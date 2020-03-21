package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.StringIsJSON

	// Comment ignored

	//lintignore:V005
	_ = v.ValidateJsonString

	_ = v.ValidateJsonString //lintignore:V005

	// Failing

	_ = v.ValidateJsonString // want "deprecated v.ValidateJsonString should be replaced with v.StringIsJSON"
}
