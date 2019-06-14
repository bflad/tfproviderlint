package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = map[string]*schema.Schema{
		"attribute_name": {
			Computed: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Optional: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Required: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Computed: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     schema.TypeString,
			},
			Computed: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Optional: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     schema.TypeString,
			},
			Optional: true,
			Type:     schema.TypeList,
		},
	}

	_ = map[string]*schema.Schema{
		"attribute_name": {
			Elem: &schema.Schema{
				Required: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     schema.TypeString,
			},
			Required: true,
			Type:     schema.TypeList,
		},
	}
}
