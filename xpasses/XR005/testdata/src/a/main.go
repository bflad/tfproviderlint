package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	// This first scenario is valid because Resource may be used within Elem
	// which should not require Description

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	_ = schema.Resource{
		Description: "test",
		Read:        readFunc,
		Schema:      map[string]*schema.Schema{},
	}

	_ = schema.Resource{
		Create:      createFunc,
		Description: "test",
		Read:        readFunc,
		Schema:      map[string]*schema.Schema{},
	}

	_ = schema.Resource{
		Create: createFunc,
		Description: "Line one.\n\n" +
			"Line two",
		Read:   readFunc,
		Schema: map[string]*schema.Schema{},
	}

	descriptions := map[string]string{"name": "test"}
	_ = schema.Resource{
		Create:      createFunc,
		Description: descriptions["name"],
		Read:        readFunc,
		Schema:      map[string]*schema.Schema{},
	}

	_ = schema.Resource{ // want "resource should configure Description"
		Read:   readFunc,
		Schema: map[string]*schema.Schema{},
	}

	_ = schema.Resource{ // want "resource should configure Description"
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*schema.Schema{},
	}

	//lintignore:XR005
	_ = schema.Resource{
		Read:   readFunc,
		Schema: map[string]*schema.Schema{},
	}

	//lintignore:XR005
	_ = schema.Resource{
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*schema.Schema{},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
