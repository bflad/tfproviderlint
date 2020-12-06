package a

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func f_v2() {
	// Comment ignored

	//lintignore:V009
	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "")

	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "") //lintignore:V009

	// Failing

	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "") // want "validation.StringMatch\\(\\) message argument should be non-empty"

	// Passing

	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "must contain only alphanumeric characters, periods, or hyphens")
}
