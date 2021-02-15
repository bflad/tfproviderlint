package a

import (
	"time"

	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{ // want "resource should not include redundant Timeouts implementation"
		Timeouts: &s.ResourceTimeout{
			Create: s.DefaultTimeout(time.Minute),
		},
	}
}
