package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Resource{Exists: existsFunc_v2} // want "resource should not include Exists function"

	_ = schema.Resource{}
}

func existsFunc_v2(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}
