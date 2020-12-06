package a

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Comment ignored

	//lintignore:V010
	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "")

	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "") //lintignore:V010

	// Failing

	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "") // want "validation.StringDoesNotMatch\\(\\) message argument should be non-empty"

	// Passing

	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "must not contain exclamation, at, octothorp, US dollar, percentage, carat, ampersand, star, or parenthesis symbols")
}
