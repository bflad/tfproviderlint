package a

import (
	"time"

	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{ // want "resource should not configure Timeouts.Create without Create implementation"
		Timeouts: &s.ResourceTimeout{
			Create: s.DefaultTimeout(time.Minute),
		},
	}

	_ = s.Resource{ // want "resource should not configure Timeouts.Delete without Delete implementation"
		Timeouts: &s.ResourceTimeout{
			Delete: s.DefaultTimeout(time.Minute),
		},
	}

	_ = s.Resource{ // want "resource should not configure Timeouts.Read without Read implementation"
		Timeouts: &s.ResourceTimeout{
			Read: s.DefaultTimeout(time.Minute),
		},
	}

	_ = s.Resource{ // want "resource should not configure Timeouts.Update without Update implementation"
		Timeouts: &s.ResourceTimeout{
			Update: s.DefaultTimeout(time.Minute),
		},
	}
}
