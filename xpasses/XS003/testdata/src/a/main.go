package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": { // want "XS003: schema might introduce crash or diff as it allows to be an empty block"
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": { // want "XS003: schema might introduce crash or diff as it allows to be an empty block"
				Required: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}

	// TypeSet block will populate its child properties even non is given
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}

	// Ignore Computed block
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
						},
					},
				},
			},
		},
	}

	// At least one of the child properties in the block is optional
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Required: true,
						},
					},
				},
			},
		},
	}

	// At least one of the child properties in the block has Default
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional: true,
							Default:  0,
						},
					},
				},
			},
		},
	}

	// At least one of the child properties in the block has DefaultFunc
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional:    true,
							DefaultFunc: func() (interface{}, error) { return nil, nil },
						},
					},
				},
			},
		},
	}

	// At least one of the child properties in the block has AtLeastOneOf constraint
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional:     true,
							AtLeastOneOf: []string{},
						},
					},
				},
			},
		},
	}

	// At least one of the child properties in the block has ExactlyOneOf constraint
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"a": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Optional:     true,
							ExactlyOneOf: []string{},
						},
					},
				},
			},
		},
	}
}
