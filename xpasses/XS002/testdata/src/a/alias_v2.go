package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{
		Schema: map[string]*s.Schema{ // want "schema attributes should be in alphabetical order"
			"name": {},
			"arn":  {},
		},
	}

	_ = s.Resource{
		Schema: map[string]*s.Schema{
			"arn": {},
			"name":  {},
		},
	}

}
