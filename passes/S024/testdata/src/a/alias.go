package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{
		Read: readFunc,
		Schema: map[string]*s.Schema{
			"test": {
				ForceNew: true, // want "ForceNew is extraneous in data source schema attributes"
				Required: true,
				Type:     s.TypeString,
			},
		},
	}
}
