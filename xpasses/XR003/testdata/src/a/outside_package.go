package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Resource{
		Create: outsideCreateFunc,
	}
}

func outsideCreateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
