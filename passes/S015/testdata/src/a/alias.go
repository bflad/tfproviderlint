package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
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
