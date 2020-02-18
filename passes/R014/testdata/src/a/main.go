package a

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Comment ignored

//lintignore:R014
func commentIgnoreFirstParameter(invalid *schema.ResourceData, meta interface{}) error {
	return nil
}

//lintignore:R014
func commentIgnoreSecondParameter(d *schema.ResourceData, invalid interface{}) error {
	return nil
}

// Failing

func failingAnonymousFirstParameter() {
	_ = func(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
		return nil
	}
}

func failingAnonymousSecondParameter() {
	_ = func(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
		return nil
	}
}

func failingFirstParameter(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
	return nil
}

func failingSecondParameter(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
	return nil
}

// Passing

func passingAnonymous() {
	_ = func(d *schema.ResourceData, meta interface{}) error {
		return nil
	}
}

func passing(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func passingOther(diff *schema.ResourceDiff, meta interface{}) error {
	return nil
}
