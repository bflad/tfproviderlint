package a

import (
	"testdata/src/a/schema"

	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func foutside() {
	_ = schema.Schema{
		Type:    s.TypeInt,
		Default: true,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:    s.TypeInt,
			Default: true,
		},
	}
}
