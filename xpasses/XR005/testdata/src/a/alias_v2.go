package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	// This first scenario is valid because Resource may be used within Elem
	// which should not require Description

	_ = s.Resource{
		Schema: map[string]*s.Schema{},
	}

	_ = s.Resource{
		Description: "test",
		Read:        readFunc_v2,
		Schema:      map[string]*s.Schema{},
	}

	_ = s.Resource{
		Create:      createFunc_v2,
		Description: "test",
		Read:        readFunc_v2,
		Schema:      map[string]*s.Schema{},
	}

	_ = s.Resource{ // want "resource should configure Description"
		Read:   readFunc_v2,
		Schema: map[string]*s.Schema{},
	}

	_ = s.Resource{ // want "resource should configure Description"
		Create: createFunc_v2,
		Read:   readFunc_v2,
		Schema: map[string]*s.Schema{},
	}

	//lintignore:XR005
	_ = s.Resource{
		Read:   readFunc_v2,
		Schema: map[string]*s.Schema{},
	}

	//lintignore:XR005
	_ = s.Resource{
		Create: createFunc_v2,
		Read:   readFunc_v2,
		Schema: map[string]*s.Schema{},
	}
}
