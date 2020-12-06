package a

import (
	"a/validation"
	"regexp"
)

func foutside() {
	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "")
}
