package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:XR006
	_ = schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	//lintignore:XR006
	_ = schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	//lintignore:XR006
	_ = schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	//lintignore:XR006
	_ = schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}
}
