package a

import (
	"a/validation"
	"regexp"
)

func foutside() {
	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "")
}
