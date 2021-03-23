package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func falias_v2() {
	_ = r.TestCase{
		IDRefreshName: "example_thing.test", // want "prefer TestStep ImportState testing over TestCase IDRefreshName"
	}
}
