package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = map[string]*schema.Resource{
		"thing": {},
	}
}
