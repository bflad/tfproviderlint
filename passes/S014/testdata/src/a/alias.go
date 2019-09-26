package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = map[string]*s.Schema{
		"attribute_name": {
			Computed: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Optional: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Required: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Type: s.TypeString,
			},
			Computed: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Type: s.TypeString,
			},
			Optional: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Type: s.TypeString,
			},
			Required: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Computed: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     s.TypeString,
			},
			Computed: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Optional: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     s.TypeString,
			},
			Optional: true,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"attribute_name": {
			Elem: &s.Schema{
				Required: true, // want "schema within Elem should not configure Computed, Optional, or Required"
				Type:     s.TypeString,
			},
			Required: true,
			Type:     s.TypeList,
		},
	}
}
