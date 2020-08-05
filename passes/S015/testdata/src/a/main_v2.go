package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = map[string]*schema.Schema{
		"INVALID": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}

	_ = map[string]*schema.Schema{
		"invalid!": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}

	_ = map[string]*schema.Schema{
		"invalid-name": {}, // want "schema attribute names should only be lowercase alphanumeric characters or underscores"
	}

	_ = map[string]*schema.Schema{
		"valid": {},
	}

	_ = map[string]*schema.Schema{
		"valid_name": {},
	}
}
