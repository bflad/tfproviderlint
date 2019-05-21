// +build tools

package main

import (
	// analysistest does not work with Go Modules yet
	_ "github.com/hashicorp/terraform/helper/schema"
	_ "github.com/hashicorp/terraform/helper/resource"
)
