package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not enable Required and Optional"
		Required: true,
		Optional: true,
	}
}
