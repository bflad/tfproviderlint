package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Resource{
		Create: createFunc_v2,
		Read:   readFunc_v2,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Create: createFunc_v2,
		Read:   readFunc_v2,
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
		Read: readFunc_v2,
		Schema: map[string]*schema.Schema{
			"test": {
				ForceNew: true, // want "ForceNew is extraneous in data source schema attributes"
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Read: readFunc_v2,
		Schema: map[string]*schema.Schema{
			"test": {
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}

	_ = schema.Resource{
		Read: readFunc_v2,
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

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}
