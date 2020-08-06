package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{Exists: existsFunc_v2} // want "resource should not include Exists function"
}
