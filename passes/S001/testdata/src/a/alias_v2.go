package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: s.TypeList,
	}

	_ = s.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: s.TypeSet,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should include Elem"
			Type: s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeList or TypeSet should include Elem"
			Type: s.TypeSet,
		},
	}
}
