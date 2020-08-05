package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// This first scenario is valid because Resource may be used within Elem
	// which should not require Description

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	_ = schema.Resource{
		Description: "test",
		Read:        readFunc_v2,
		Schema:      map[string]*schema.Schema{},
	}

	_ = schema.Resource{
		Create:      createFunc_v2,
		Description: "test",
		Read:        readFunc_v2,
		Schema:      map[string]*schema.Schema{},
	}

	_ = schema.Resource{ // want "resource should configure Description"
		Read:   readFunc_v2,
		Schema: map[string]*schema.Schema{},
	}

	_ = schema.Resource{ // want "resource should configure Description"
		Create: createFunc_v2,
		Read:   readFunc_v2,
		Schema: map[string]*schema.Schema{},
	}

	//lintignore:XR005
	_ = schema.Resource{
		Read:   readFunc_v2,
		Schema: map[string]*schema.Schema{},
	}

	//lintignore:XR005
	_ = schema.Resource{
		Create: createFunc_v2,
		Read:   readFunc_v2,
		Schema: map[string]*schema.Schema{},
	}
}

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}
