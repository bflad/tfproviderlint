package a

import (
	"a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5")
}
