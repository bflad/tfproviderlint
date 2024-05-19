package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.GetOkExists("test")
}
