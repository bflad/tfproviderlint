package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		Read: outsideReadFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func outsideReadFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
