package a

import (
	"a/schema"
	"time"
)

func foutside() {
	oneMinute := time.Minute
	_ = schema.Resource{
		Create: outsideCreateFunc,
		Timeouts: &schema.ResourceTimeout{
			Create: &oneMinute,
		},
	}
}

func outsideCreateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
