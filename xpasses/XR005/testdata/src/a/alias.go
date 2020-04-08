package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	// This first scenario is valid because Resource may be used within Elem
	// which should not require Description

	_ = s.Resource{
		Schema: map[string]*s.Schema{},
	}

	_ = s.Resource{
		Description: "test",
		Read:        readFunc,
		Schema:      map[string]*s.Schema{},
	}

	_ = s.Resource{
		Create:      createFunc,
		Description: "test",
		Read:        readFunc,
		Schema:      map[string]*s.Schema{},
	}

	_ = s.Resource{ // want "resource should configure Description"
		Read:   readFunc,
		Schema: map[string]*s.Schema{},
	}

	_ = s.Resource{ // want "resource should configure Description"
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*s.Schema{},
	}

	//lintignore:XR005
	_ = s.Resource{
		Read:   readFunc,
		Schema: map[string]*s.Schema{},
	}

	//lintignore:XR005
	_ = s.Resource{
		Create: createFunc,
		Read:   readFunc,
		Schema: map[string]*s.Schema{},
	}
}
