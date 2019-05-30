package a

import (
	"a/schema"
)

func foutside() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "")
}
