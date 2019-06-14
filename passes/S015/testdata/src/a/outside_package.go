package a

import (
	"a/schema"
)

func foutside() {
	_ = map[string]*schema.Schema{
		"INVALID": {},
	}

	_ = map[string]*schema.Schema{
		"invalid!": {},
	}

	_ = map[string]*schema.Schema{
		"invalid-name": {},
	}
}
