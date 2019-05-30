package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not enable Required and Computed"
		Required: true,
		Computed: true,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = schema.Schema{
		Computed: true,
	}
}
