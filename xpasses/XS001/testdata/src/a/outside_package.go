package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = map[string]*schema.Schema{
		"name": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
