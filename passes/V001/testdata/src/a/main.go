package a

import (
	"errors"
	"regexp"
)

//lintignore:V001
func commentIgnore(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)

	if !regexp.MustCompile(`example`).MatchString(value) {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringMatch\\(\\) or validation.StringDoesNotMatch\\(\\)"
		value := v.(string)

		if !regexp.MustCompile(`example`).MatchString(value) {
			errs = append(errs, errors.New("example"))
		}

		return
	}
}

func failingInlineRegexp(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringMatch\\(\\) or validation.StringDoesNotMatch\\(\\)"
	value := v.(string)

	if !regexp.MustCompile(`example`).MatchString(value) {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingVariableRegexp(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringMatch\\(\\) or validation.StringDoesNotMatch\\(\\)"
	value := v.(string)

	re := regexp.MustCompile(`example`)

	if !re.MatchString(value) {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) {
		return
	}
}

func passingRegexpMatchString() {
	_ = regexp.MustCompile(`example`).MatchString("")
}

func passingSchemaValidateFunc(v interface{}, k string) (warns []string, errs []error) {
	return
}
