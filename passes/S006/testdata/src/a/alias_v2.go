package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema of TypeMap should include Elem"
		Type: s.TypeMap,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: s.TypeMap,
		},
	}
}
