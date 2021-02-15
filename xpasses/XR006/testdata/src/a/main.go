package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{ // want "resource should not include redundant Timeouts implementation"
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{ // want "resource should not include redundant Timeouts implementation"
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}


	_ = schema.Resource{
		Create:   createFunc,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}