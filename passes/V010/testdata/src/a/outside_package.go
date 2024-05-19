package a

import (
	"regexp"
	"testdata/src/a/validation"
)

func foutside() {
	validation.StringDoesNotMatch(regexp.MustCompile(`^[!@#$%^&*()]+$`), "")
}
