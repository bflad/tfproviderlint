package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{Exists: existsFunc} // want "resource should not include Exists function"

	_ = schema.Resource{}
}

func existsFunc(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}
