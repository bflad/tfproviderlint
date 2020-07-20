package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {}, // want "schema attribute name should be in alphabetical order"
			"arn": {},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"arn": {},
			"valid":  {},
		},
	}
}



