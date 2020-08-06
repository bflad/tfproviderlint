package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = map[string]*s.Schema{
		"INVALID": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}

	_ = map[string]*s.Schema{
		"invalid!": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}

	_ = map[string]*s.Schema{
		"invalid-name": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}
}
