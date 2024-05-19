package a

import (
	"testdata/src/a/schema"
	"testdata/src/a/time"
)

func foutside() {
	var d schema.ResourceData

	d.SetId(time.Now())
}
