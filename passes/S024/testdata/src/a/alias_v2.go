package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{
		Read: readFunc_v2,
		Schema: map[string]*s.Schema{
			"test": {
				ForceNew: true, // want "ForceNew is extraneous in data source schema attributes"
				Required: true,
				Type:     s.TypeString,
			},
		},
	}
}
