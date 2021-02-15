package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
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
