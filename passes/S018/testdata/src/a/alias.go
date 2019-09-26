package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should use TypeList with MaxItems 1"
		MaxItems: 1,
		Type:     s.TypeSet,
	}

	_ = s.Schema{
		MaxItems: 2,
		Type:     s.TypeSet,
	}

	_ = s.Schema{
		Type: s.TypeSet,
	}

	_ = s.Schema{
		MaxItems: 1,
		Type:     s.TypeList,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should use TypeList with MaxItems 1"
			MaxItems: 1,
			Type:     s.TypeSet,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MaxItems: 2,
			Type:     s.TypeSet,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeSet,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MaxItems: 1,
			Type:     s.TypeList,
		},
	}
}
