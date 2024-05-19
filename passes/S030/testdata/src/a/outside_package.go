package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:     true,
		InputDefault: "test",
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:     true,
			InputDefault: "test",
		},
	}
}
