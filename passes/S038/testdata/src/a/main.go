package a

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func f() {
	// Comment ignored

	//lintignore:S038
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							AtLeastOneOf: []string{"x.0.bar"},
						},
						"bar": {
							AtLeastOneOf: []string{"x.0.foo"},
						},
					},
				},
			},
		},
	}
	// Failing

	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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

	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							AtLeastOneOf: []string{"x.1.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.1" configuration block attribute references must be separated by .0`
						},
						"bar": {
							AtLeastOneOf: []string{"x.1.foo"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.1" configuration block attribute references must be separated by .0`
						},
					},
				},
			},
		},
	}

	// Passing

	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							AtLeastOneOf: []string{"x.0.bar"},
						},
						"bar": {
							AtLeastOneOf: []string{"x.0.foo"},
						},
					},
				},
			},
		},
	}

}
