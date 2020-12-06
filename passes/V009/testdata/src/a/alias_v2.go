package a

import (
	"regexp"

	v "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func falias_v2() {
	// Comment ignored

	//lintignore:V009
	v.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "")

	v.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "") //lintignore:V009

	// Failing

	v.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "") // want "validation.StringMatch\\(\\) message argument should be non-empty"

	// Passing

	v.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "must contain only alphanumeric characters, periods, or hyphens")
}
