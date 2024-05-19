package a

import (
	"testdata/src/a/resource"
)

func foutside() {
	_ = resource.TestCase{
		IDRefreshName: "example_thing.test",
	}
}
