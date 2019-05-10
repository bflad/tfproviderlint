package a

import (
	r "github.com/hashicorp/terraform/helper/resource"
)

func falias() {
	_ = r.TestCase{} // want "missing CheckDestroy"
}
