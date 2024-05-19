package a

import (
	"testdata/src/a/resource"
	"testdata/src/a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.SetId(resource.PrefixedUniqueId("test"))
}
