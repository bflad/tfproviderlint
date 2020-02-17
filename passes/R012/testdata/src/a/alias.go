package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{ // want "data source should not configure CustomizeDiff"
		CustomizeDiff: customizeDiffFunc,
		Read:          readFunc,
	}
}
