package a

import (
	"errors"
	"strconv"
)

//lintignore:V012
func commentIgnore(v interface{}, k string) (warns []string, errs []error) {
	value := v.(int)

	if value > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
		value := v.(int)

		if value > 255 {
			errs = append(errs, errors.New("example"))
		}

		return
	}
}

func failingSchemaValidateFuncGreaterThan(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
	value := v.(int)

	if value > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncGreaterThanOrEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
	value := v.(int)

	if value >= 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncLessThan(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
	value := v.(int)

	if value < 3 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncLessThanOrEqual(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
	value := v.(int)

	if value <= 3 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func failingSchemaValidateFuncMultiple(v interface{}, k string) (warns []string, errs []error) { // want "custom SchemaValidateFunc should be replaced with validation.IntAtLeast\\(\\), validation.IntAtMost\\(\\), or validation.IntBetween\\(\\)"
	value := v.(int)

	if value < 3 || value > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingAnonymousFunction() {
	_ = func(v interface{}, k string) (warns []string, errs []error) {
		return
	}
}

func passingIntCheck() {
	var value int
	_ = value > 255
}

func passingSchemaValidateFunc(v interface{}, k string) (warns []string, errs []error) {
	return
}

func passingSchemaValidateFuncString(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)

	if len(value) > 255 {
		errs = append(errs, errors.New("example"))
	}

	return
}

func passingSchemaValidateFuncStrconvAtoi(v interface{}, k string) (warns []string, errs []error) {
	vInt, _ := strconv.Atoi(v.(string))

	if vInt < 0 || vInt > 100 {
		errs = append(errs, errors.New("example"))
	}

	return
}
