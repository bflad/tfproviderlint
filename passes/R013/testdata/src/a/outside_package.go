package a

import (
	"a/schema"
)

func foutside() {
	_ = map[string]*schema.Resource{
		"thing": {},
	}
}
