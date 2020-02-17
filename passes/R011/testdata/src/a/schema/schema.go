package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type Resource struct {
	MigrateState  func(int, *terraform.InstanceState, interface{}) (*terraform.InstanceState, error)
	SchemaVersion int
}
