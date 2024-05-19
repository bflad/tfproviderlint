package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Set:  func() {},
		Type: schema.TypeList,
	}

	_ = map[string]*schema.Schema{
		"name": {
			Set:  func() {},
			Type: schema.TypeList,
		},
	}
}
