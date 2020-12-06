package validation

import "regexp"

type SchemaValidateFunc func(interface{}, string) ([]string, []error)

func StringMatch(r *regexp.Regexp, message string) SchemaValidateFunc {
	return nil
}
