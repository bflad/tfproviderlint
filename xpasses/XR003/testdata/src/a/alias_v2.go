package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{ // want "resource should include Timeouts implementation"
		Create: createFunc_v2,
	}
}
