package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = &s.Resource{
		Read: func(*s.ResourceData, interface{}) error { return nil },
		Schema: map[string]*s.Schema{
			"x": {
				Type: s.TypeList,
				Elem: &s.Resource{
					Schema: map[string]*s.Schema{
						"foo": {
							AtLeastOneOf: []string{"x.0.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.0" configuration block attribute references are only valid for TypeList and MaxItems: 1 attributes`
						},
						"bar": {
							AtLeastOneOf: []string{"x.0.foo"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.0" configuration block attribute references are only valid for TypeList and MaxItems: 1 attributes`
						},
					},
				},
			},
		},
	}
}
