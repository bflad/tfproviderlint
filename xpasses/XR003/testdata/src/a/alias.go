package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{ // want "resource should include Timeouts implementation"
		Create: createFunc,
	}
}
