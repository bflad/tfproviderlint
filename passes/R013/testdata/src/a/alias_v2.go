package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = map[string]*s.Resource{
		"thing": {}, // want "resource names should include the provider name and at least one underscore"
	}
}
