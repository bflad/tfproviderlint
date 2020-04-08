package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		Read:   outsideReadFunc,
		Schema: map[string]*schema.Schema{},
	}
}

func outsideReadFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
