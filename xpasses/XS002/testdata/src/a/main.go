package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {}, // want "schema attribute name should be in alphabetical order"
			"arn":  {}, // want "schema attribute name should be in alphabetical order"
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {}, // want "schema attribute name should be in alphabetical order"
			"arn":  {}, // want "schema attribute name should be in alphabetical order"
			"block": {}, // want "schema attribute name should be in alphabetical order"
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"block": {}, // want "schema attribute name should be in alphabetical order"
			"arn":  {}, // want "schema attribute name should be in alphabetical order"
			"name": {},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn":  {},
			"name": {}, // want "schema attribute name should be in alphabetical order"
			"block": {}, // want "schema attribute name should be in alphabetical order"
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn": {},
			"name":  {},
		},
	}
}
