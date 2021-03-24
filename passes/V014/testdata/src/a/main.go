package a

import (
	"errors"
	"strconv"
)

//lintignore:V014
func commentIgnore(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value != 1 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntInSlice\\(\\) or validation.IntNotInSlice\\(\\)"
		value := v.(int)

		if value != 1 {
			errs = append(errs, errors.New("example"))
		}

		return
	}
}

func failingSchemaValidateFuncNotEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntInSlice\\(\\) or validation.IntNotInSlice\\(\\)"
	value := v.(int)

	if value != 1 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntInSlice\\(\\) or validation.IntNotInSlice\\(\\)"
	value := v.(int)

	if value == 1 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncMultiple(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntInSlice\\(\\) or validation.IntNotInSlice\\(\\)"
	value := v.(int)

	if value != 1 && value != 2 {
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
	var value int
	_ = value != 1
}

func passingSchemaValidateFunc(v interface{}, k string) (warns []string, errs []error) {
	return
}

func passingSchemaValidateFuncIntGreaterThan(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncIntGreaterThanOrEqual(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value >= 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncIntLessThan(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value < 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncIntLessThanOrEqual(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value <= 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncStrconvAtoi(v interface{}, k string) (warns []string, errs []error) {
	vInt, _ := strconv.Atoi(v.(string))

	if vInt != 1 && vInt != 2 {
		errs = append(errs, errors.New("example"))
	}

	return
}
