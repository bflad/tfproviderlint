package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
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
		Create:   createFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}
}

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}