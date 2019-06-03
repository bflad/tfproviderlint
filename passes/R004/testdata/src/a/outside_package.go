package a

import (
	"time"

	"a/schema"
)

func foutside() {
	var d schema.ResourceData
	var t time.Time

	d.Set("test", t)
	d.Set("test", []time.Time{t})
	d.Set("test", map[string]time.Time{"test": t})
}
