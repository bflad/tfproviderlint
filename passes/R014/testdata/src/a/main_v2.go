package a

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Comment ignored

//lintignore:R014
func commentIgnoreFirstParameter_v2(invalid *schema.ResourceData, meta interface{}) error {
	return nil
}

//lintignore:R014
func commentIgnoreSecondParameter_v2(d *schema.ResourceData, invalid interface{}) error {
	return nil
}

// Failing

func failingAnonymousFirstParameter_v2() {
	_ = func(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
		return nil
	}
}

func failingAnonymousSecondParameter_v2() {
	_ = func(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
		return nil
	}
}

func failingFirstParameter_v2(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
	return nil
}

func failingSecondParameter_v2(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
	return nil
}

// Passing

func passingAnonymous_v2() {
	_ = func(d *schema.ResourceData, meta interface{}) error {
		return nil
	}
}

func passing_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func passingOther_v2(diff *schema.ResourceDiff, meta interface{}) error {
	return nil
}
