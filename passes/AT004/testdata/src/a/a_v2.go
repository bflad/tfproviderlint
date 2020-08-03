package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func f_v2() {
	var t *testing.T

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `provider "aws" {}`, // want "provider declaration should be omitted"
			},
		},
	}

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: ``,
			},
		},
	}

	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `provider "aws" {}`, // want "provider declaration should be omitted"
			},
		},
	})

	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: ``,
			},
		},
	})
}
