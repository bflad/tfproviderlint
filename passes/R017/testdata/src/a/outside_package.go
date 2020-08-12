package a

import (
	"a/schema"
	"a/time"
)

func foutside() {
	var d schema.ResourceData

	d.SetId(time.Now())
}
