package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not enable Required and configure Default"
		Required: true,
		Default:  true,
	}

	_ = schema.Schema{
		Required: true,
		Default:  nil,
	}

	_ = schema.Schema{
		Required: true,
	}

	_ = schema.Schema{
		Default: true,
	}
}
