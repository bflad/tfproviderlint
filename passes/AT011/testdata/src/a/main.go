package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func f() {
	// Passing

	_ = resource.TestCase{}

	_ = resource.TestCase{
		IDRefreshName:   "example_thing.test",
		IDRefreshIgnore: []string{"attr1"},
	}

	// Comment ignored

	_ = resource.TestCase{
		//lintignore:AT011
		IDRefreshIgnore: []string{"attr1"},
	}

	// Failing

	_ = resource.TestCase{
		IDRefreshIgnore: []string{"attr1"}, // want "extraneous TestCase IDRefreshIgnore without IDRefreshName"
	}
}
