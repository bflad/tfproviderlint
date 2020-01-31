package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"test": {
							ForceNew: true,
							Required: true,
							Type:     schema.TypeString,
						},
					},
				},
				Required: true,
				Type:     schema.TypeList,
			},
		},
	}

	_ = schema.Resource{
		Read: readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true, // want "ForceNew is extraneous in data source schema attributes"
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Read: readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Read:   readFunc,
		Schema: map[string]*schema.Schema{
			"test": {
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"test": {
							ForceNew: true, // want "ForceNew is extraneous in data source schema attributes"
							Required: true,
							Type:     schema.TypeString,
						},
					},
				},
				Required: true,
				Type:     schema.TypeList,
			},
		},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
