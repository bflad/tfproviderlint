package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.StringIsValidRegExp

	// Comment ignored

	//lintignore:V007
	_ = validation.ValidateRegexp

	_ = validation.ValidateRegexp //lintignore:V007

	// Failing

	_ = validation.StringIsValidRegExp // want "deprecated validation.ValidateRegexp should be replaced with validation.StringIsValidRegExp"
}
