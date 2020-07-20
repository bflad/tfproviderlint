package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{
		Schema: map[string]*s.Schema{
			"name": {}, // want "schema attribute name should be in alphabetical order"
			"arn":  {}, // want "schema attribute name should be in alphabetical order"
		},
	}

	_ = s.Resource{
		Schema: map[string]*s.Schema{
			"arn": {},
			"name":  {},
		},
	}

}
