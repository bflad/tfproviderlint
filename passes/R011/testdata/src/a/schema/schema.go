package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type Resource struct {
	MigrateState  func(int, *terraform.InstanceState, interface{}) (*terraform.InstanceState, error)
	SchemaVersion int
}
