package a

import (
	"a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.GetOkExists("test")
}
