package a

import (
	"errors"
)

//lintignore:V011
func commentIgnore(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)

	if len(value) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
		value := v.(string)

		if len(value) > 255 {
			errs = append(errs, errors.New("example"))
		}

		return
	}
}

func failingSchemaValidateFuncGreaterThan(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
	value := v.(string)

	if len(value) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncGreaterThanOrEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
	value := v.(string)

	if len(value) >= 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncLessThan(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
	value := v.(string)

	if len(value) < 3 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncLessThanOrEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
	value := v.(string)

	if len(value) <= 3 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncMultiple(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringLenBetween\\(\\)"
	value := v.(string)

	if len(value) < 3 || len(value) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) {
		return
	}
}

func passingLen() {
	var value string
	_ = len(value) > 255
}

func passingSchemaValidateFunc(v interface{}, k string) (warns []string, errs []error) {
	return
}

func passingSchemaValidateFuncNonString1(v interface{}, k string) (warns []string, errs []error) {
	var test []string

	if len(test) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncNonString2(v interface{}, k string) (warns []string, errs []error) {
	var test map[string]string

	if len(test) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}
