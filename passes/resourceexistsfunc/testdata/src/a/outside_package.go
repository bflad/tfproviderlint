package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{Exists: outsideExistsFunc}
}

func outsideExistsFunc(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}
