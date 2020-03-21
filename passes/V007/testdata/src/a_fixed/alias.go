package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.StringIsValidRegExp

	// Comment ignored

	//lintignore:V007
	_ = v.ValidateRegexp

	_ = v.ValidateRegexp //lintignore:V007

	// Failing

	_ = v.StringIsValidRegExp // want "deprecated v.ValidateRegexp should be replaced with v.StringIsValidRegExp"
}
