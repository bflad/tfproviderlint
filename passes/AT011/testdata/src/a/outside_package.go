package a

import (
	"a/resource"
)

func foutside() {
	_ = resource.TestCase{
		IDRefreshIgnore: []string{"attr1"},
	}
}
