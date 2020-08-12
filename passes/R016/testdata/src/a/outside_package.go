package a

import (
	"a/resource"
	"a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.SetId(resource.PrefixedUniqueId("test"))
}
