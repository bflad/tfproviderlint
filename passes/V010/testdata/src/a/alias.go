package a

import (
	"regexp"

	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Comment ignored

	//lintignore:V010
	v.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "")

	v.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "") //lintignore:V010

	// Failing

	v.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "") // want "validation.StringDoesNotMatch\\(\\) message argument should be non-empty"

	// Passing

	v.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "must not contain exclamation, at, octothorp, US dollar, percentage, carat, ampersand, star, or parenthesis symbols")
}
