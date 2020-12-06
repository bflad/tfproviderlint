package validation

import "regexp"

type SchemaValidateFunc func(interface{}, string) ([]string, []error)

func StringDoesNotMatch(r *regexp.Regexp, message string) SchemaValidateFunc {
	return nil
}
