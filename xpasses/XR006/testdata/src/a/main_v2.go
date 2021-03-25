package a

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
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
		Create: createFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Delete: deleteFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Read: readFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		Update: updateFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		CreateContext: createContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		DeleteContext: deleteContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		ReadContext: readContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		UpdateContext: updateContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		CreateWithoutTimeout: createContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		DeleteWithoutTimeout: deleteContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		ReadWithoutTimeout: readContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(time.Minute),
		},
	}

	_ = schema.Resource{
		UpdateWithoutTimeout: updateContextFunc_v2,
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(time.Minute),
		},
	}
}

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func createContextFunc_v2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func deleteFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteContextFunc_v2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func readFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readContextFunc_v2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func updateFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateContextFunc_v2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
