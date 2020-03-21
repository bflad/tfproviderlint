package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.IsRFC3339Time

	// Comment ignored

	//lintignore:V008
	_ = v.ValidateRFC3339TimeString

	_ = v.ValidateRFC3339TimeString //lintignore:V008

	// Failing

	_ = v.IsRFC3339Time // want "deprecated v.ValidateRFC3339TimeString should be replaced with v.IsRFC3339Time"
}
