package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{
		Type: s.TypeMap,
		Elem: &s.Resource{}, // want "schema of TypeMap should not use Elem of \\*schema.Resource"
	}

	_ = s.Schema{
		Type: s.TypeMap,
		Elem: &s.Schema{},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeMap,
			Elem: &s.Resource{}, // want "schema of TypeMap should not use Elem of \\*schema.Resource"
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeMap,
			Elem: &s.Schema{},
		},
	}
}
