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

	// Configuration block has not specified `MaxItems: 1`
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

	// Configuration block' type is not `TypeList`
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeSet,
				MaxItems: 1,
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

	// Invalid reference: reference inside configuration block (not 0th block)
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
					},
				},
			},
		},
	}

	// Invalid reference: reference inside configuration block (non-existed attribute)
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							AtLeastOneOf: []string{"x.0.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.0.bar" references to unknown attribute`
						},
					},
				},
			},
		},
	}

	// Invalid reference: reference inside configuration block (deep)
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"y": {
							Type: schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"foo": {
										AtLeastOneOf: []string{"x.0.y.0.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.0.y.0.bar" references to unknown attribute`
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Invalid reference: root attribute reference non-existing root reference
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "bar" references to unknown attribute`
			},
		},
	}

	// Invalid reference: root attribute referencing existing configuration block attribute with non-existing nested attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"x.0.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "x.0.bar" references to unknown attribute`
			},
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}

	// Invalid reference: configuration block nested attribute referencing non-existing root attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Type:         schema.TypeString,
							AtLeastOneOf: []string{"bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "bar" references to unknown attribute`
						},
					},
				},
			},
		},
	}

	// Invalid reference: configuration block nested attribute referencing existing configuration block attribute with non-existing nested attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Type:         schema.TypeString,
							AtLeastOneOf: []string{"y.0.bar"}, // want `S038: invalid AtLeastOneOf attribute reference semantics: "y.0.bar" references to unknown attribute`
						},
					},
				},
			},
			"y": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}

	// Passing

	// Valid reference: reference inside configuration block
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

	// Valid reference: reference inside configuration block (deep)
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"y": {
							Type: schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"foo": {
										AtLeastOneOf: []string{"x.0.y.0.bar"},
									},
									"bar": {
										AtLeastOneOf: []string{"x.0.y.0.foo"},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Valid reference: root attribute reference root reference
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"bar"},
			},
			"bar": {
				Type: schema.TypeString,
			},
		},
	}

	// Valid reference: root attribute referencing existing configuration block attribute with existing nested attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"x.0.bar"},
			},
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bar": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}

	// Valid reference: configuration block nested attribute referencing existing root attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Type:         schema.TypeString,
							AtLeastOneOf: []string{"bar"},
						},
					},
				},
			},
			"bar": {
				Type: schema.TypeString,
			},
		},
	}

	// Valid reference: configuration block nested attribute referencing existing configuration block attribute with existing nested attribute
	_ = &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error { return nil },
		Schema: map[string]*schema.Schema{
			"x": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Type:         schema.TypeString,
							AtLeastOneOf: []string{"y.0.bar"},
						},
					},
				},
			},
			"y": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bar": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}

}
