package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include Default"
		Default: true,
		Type:    s.TypeList,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should not include Default"
		Default: true,
		Type:    s.TypeSet,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include Default"
			Default: true,
			Type:    s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include Default"
			Default: true,
			Type:    s.TypeSet,
		},
	}
}
