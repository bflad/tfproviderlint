package a

import (
	"a/schema"
)

func foutside() {
	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			},
			Computed: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			Optional: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			Required: true,
			Type:     schema.TypeList,
		},
	}
}
