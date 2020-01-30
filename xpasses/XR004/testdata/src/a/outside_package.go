package a

import (
	"a/schema"
)

func foutside() {
	var d schema.ResourceData

	d.Set("test", []interface{}{})
	d.Set("test", map[string]interface{}{})
}
