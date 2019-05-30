package a

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func f() {
	var t *testing.T

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `provider "aws" {}`, // want "provider declaration should be omitted"
			},
		},
	}

	/*
		Skipping for now since there does not seem to be a way to
		correctly add the want comment for multiline string literals

		_ = resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: configWithProvider1,
				},
			},
		}

		_ = resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: configWithProvider2,
				},
			},
		}

		_ = resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: configWithProvider3(),
				},
			},
		}
	*/

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: ``,
			},
		},
	}

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: configWithoutProvider1,
			},
		},
	}

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: configWithoutProvider2,
			},
		},
	}

	_ = resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: configWithoutProvider3(),
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

/*
Skipping for now since there does not seem to be a way to
correctly add the want comment for multiline string literals

const configWithProvider1 = `
provider "aws" {
  region = "us-east-1"
}
`

var configWithProvider2 = `
provider "aws" {
  region = "us-east-1"
}
`

func configWithProvider3() string {
	return fmt.Sprintf(`
provider "aws" {
  region = "us-east-1"
}
`)
}
*/

const configWithoutProvider1 = `
resource "aws_vpc" {
  cidr_block = "10.0.0.0/16"
}
`

var configWithoutProvider2 = `
resource "aws_vpc" {
  cidr_block = "10.0.0.0/16"
}
`

func configWithoutProvider3() string {
	return fmt.Sprintf(`
resource "aws_vpc" {
  cidr_block = "10.0.0.0/16"
}
`)
}
