package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{ // want "resource should include Importer implementation"
		Create: createFunc,
	}

	_ = schema.Resource{
		Read: readFunc,
	}

	_ = schema.Resource{
		Create:   createFunc,
		Importer: &schema.ResourceImporter{},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
