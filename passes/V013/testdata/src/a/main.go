package a

import (
	"errors"
)

//lintignore:V013
func commentIgnore(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)

	if value != "test" {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringInSlice\\(\\) or validation.StringNotInSlice\\(\\)"
		value := v.(string)

		if value != "test" {
			errs = append(errs, errors.New("example"))
		}

		return
	}
}

func failingSchemaValidateFuncNotEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringInSlice\\(\\) or validation.StringNotInSlice\\(\\)"
	value := v.(string)

	if value != "test" {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringInSlice\\(\\) or validation.StringNotInSlice\\(\\)"
	value := v.(string)

	if value == "test" {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncMultiple(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.StringInSlice\\(\\) or validation.StringNotInSlice\\(\\)"
	value := v.(string)

	if value != "test1" && value != "test2" {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) {
		return
	}
}

func passingStringNotEqual() {
	var value string
	_ = value != "test"
}

func passingSchemaValidateFunc(v interface{}, k string) (warns []string, errs []error) {
	return
}

func passingSchemaValidateFuncStringLen(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)

	if len(value) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}
