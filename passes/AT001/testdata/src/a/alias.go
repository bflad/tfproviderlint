package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func falias() {
	_ = r.TestCase{} // want "missing CheckDestroy"
}
