package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func falias_v2() {
	_ = r.TestCase{} // want "missing ErrorCheck"
}
