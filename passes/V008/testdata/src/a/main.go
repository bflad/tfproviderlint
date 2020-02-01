package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.IsRFC3339Time

	// Comment ignored

	//lintignore:V008
	_ = validation.ValidateRFC3339TimeString

	_ = validation.ValidateRFC3339TimeString //lintignore:V008

	// Failing

	_ = validation.ValidateRFC3339TimeString // want "deprecated validation.ValidateRFC3339TimeString should be replaced with validation.IsRFC3339Time"
}
