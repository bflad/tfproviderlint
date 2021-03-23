package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func falias_v2() {
	_ = r.TestCase{
		IDRefreshIgnore: []string{"attr1"}, // want "extraneous TestCase IDRefreshIgnore without IDRefreshName"
	}
}
