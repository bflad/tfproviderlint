package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func f() {
	// Passing

	_ = resource.TestCase{}

	// Comment ignored

	_ = resource.TestCase{
		//lintignore:AT010
		IDRefreshName: "example_thing.test",
	}

	// Failing

	_ = resource.TestCase{
		IDRefreshName: "example_thing.test", // want "prefer TestStep ImportState testing over TestCase IDRefreshName"
	}
}
