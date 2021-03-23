package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func falias() {
	_ = r.TestCase{
		IDRefreshName: "example_thing.test", // want "prefer TestStep ImportState testing over TestCase IDRefreshName"
	}
}
