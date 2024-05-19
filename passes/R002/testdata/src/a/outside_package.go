package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	var d schema.ResourceData
	var stringPtr *string

	d.Set("test", *stringPtr)
}
