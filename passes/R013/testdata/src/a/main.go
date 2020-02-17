package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	// Comment ignored

	//lintignore:R013
	_ = map[string]*schema.Resource{
		"thing": {},
	}

	// Failing

	_ = map[string]*schema.Resource{
		"thing": {}, // want "resource names should include the provider name and at least one underscore"
	}

	// Passing

	_ = map[string]*schema.Resource{
		"example_thing": {},
	}
}
