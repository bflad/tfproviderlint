package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{ // want "resource should not configure Timeouts.Create without Create implementation"
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{ // want "resource should not configure Timeouts.Delete without Delete implementation"
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{ // want "resource should not configure Timeouts.Read without Read implementation"
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{ // want "resource should not configure Timeouts.Update without Update implementation"
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Create: createFunc,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Delete: deleteFunc,
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Read: readFunc,
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Update: updateFunc,
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
