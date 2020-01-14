package a

import (
	"a/schema"
)

func foutside() {
	var d schema.ResourceData
	
	_ = d.HasChange("attr1") || d.HasChange("attr2")

	_ = d.HasChange("attr1") ||
		d.HasChange("attr2")
}
