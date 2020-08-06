package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Resource{ // want "resource should include Importer implementation"
		Create: createFunc_v2,
	}

	_ = schema.Resource{
		Read: readFunc_v2,
	}

	_ = schema.Resource{
		Create:   createFunc_v2,
		Importer: &schema.ResourceImporter{},
	}
}

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}
