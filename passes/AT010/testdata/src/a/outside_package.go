package a

import (
	"a/resource"
)

func foutside() {
	_ = resource.TestCase{
		IDRefreshName: "example_thing.test",
	}
}
