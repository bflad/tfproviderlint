package a

import (
	"testdata/src/a/resource"
)

func foutside() {
	_ = resource.TestCase{
		IDRefreshIgnore: []string{"attr1"},
	}
}
