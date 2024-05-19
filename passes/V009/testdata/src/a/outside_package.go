package a

import (
	"regexp"
	"testdata/src/a/validation"
)

func foutside() {
	validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9.-]+$`), "")
}
