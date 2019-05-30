package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Required: true,
		Optional: true,
	}
}
