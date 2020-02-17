package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = map[string]*s.Resource{
		"thing": {}, // want "resource names should include the provider name and at least one underscore"
	}
}
