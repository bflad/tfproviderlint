package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func falias() {
	_ = r.TestCase{
		IDRefreshIgnore: []string{"attr1"}, // want "extraneous TestCase IDRefreshIgnore without IDRefreshName"
	}
}
