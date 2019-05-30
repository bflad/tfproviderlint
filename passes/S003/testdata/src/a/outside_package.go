package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required: true,
		Computed: true,
	}
}
