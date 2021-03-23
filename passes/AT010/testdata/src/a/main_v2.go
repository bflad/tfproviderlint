package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func f_v2() {
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
