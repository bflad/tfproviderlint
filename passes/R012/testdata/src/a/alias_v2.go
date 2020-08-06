package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{ // want "data source should not configure CustomizeDiff"
		CustomizeDiff: customizeDiffFunc_v2,
		Read:          readFunc_v2,
	}
}
