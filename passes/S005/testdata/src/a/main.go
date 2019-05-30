package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not enable Computed and configure Default"
		Computed: true,
		Default:  true,
	}

	_ = schema.Schema{
		Computed: true,
		Default:  nil,
	}

	_ = schema.Schema{
		Computed: true,
	}

	_ = schema.Schema{
		Default: true,
	}
}
