package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
		MaxItems: 1,
		Type:     s.TypeString,
	}

	_ = s.Schema{ // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
		MinItems: 1,
		Type:     s.TypeString,
	}

	_ = s.Schema{
		Type: s.TypeString,
	}

	_ = s.Schema{
		MaxItems: 1,
		Type:     s.TypeList,
	}

	_ = s.Schema{
		MaxItems: 1,
		Type:     s.TypeMap,
	}

	_ = s.Schema{
		MaxItems: 1,
		Type:     s.TypeSet,
	}

	_ = s.Schema{
		MinItems: 1,
		Type:     s.TypeList,
	}

	_ = s.Schema{
		MinItems: 1,
		Type:     s.TypeMap,
	}

	_ = s.Schema{
		MinItems: 1,
		Type:     s.TypeSet,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
			MaxItems: 1,
			Type:     s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema MaxItems or MinItems should only be included for TypeList, TypeMap, or TypeSet"
			MinItems: 1,
			Type:     s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MaxItems: 1,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MaxItems: 1,
			Type:     s.TypeMap,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MaxItems: 1,
			Type:     s.TypeSet,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MinItems: 1,
			Type:     s.TypeList,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MinItems: 1,
			Type:     s.TypeMap,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			MinItems: 1,
			Type:     s.TypeSet,
		},
	}
}
