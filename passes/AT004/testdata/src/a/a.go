package a

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
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

const configWithProvider1 = /* want "provider declaration should be omitted" */ `
provider "aws" {
  region = "us-east-1"
}
`

var configWithProvider2 = /* want "provider declaration should be omitted" */ `
provider "aws" {
  region = "us-east-1"
}
`

func configWithProvider3() string {
	return fmt.Sprintf( /* want "provider declaration should be omitted" */ `
provider "aws" {
  region = "us-east-1"
}
`)
}

//lintignore:AT004
const configWithProviderIgnored1 = `
provider "aws" {
  region = "us-east-1"
}
`

//lintignore:AT004
var configWithProviderIgnored2 = `
provider "aws" {
  region = "us-east-1"
}
`

func configWithProviderIgnored3() string {
	//lintignore:AT004
	return fmt.Sprintf(`
provider "aws" {
  region = "us-east-1"
}
`)
}

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
