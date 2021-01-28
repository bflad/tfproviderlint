package a

import (
	"time"

	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{ // want "resource should not include redundant Timeouts implementation"
		Timeouts: &s.ResourceTimeout{
			Create: s.DefaultTimeout(time.Minute),
		},
	}
}
