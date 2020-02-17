package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func f() {
	// Comment ignored

	//lintignore:R011
	_ = schema.Resource{
		MigrateState:  migrateStateFunc,
		SchemaVersion: 1,
	}

	// Failing

	_ = schema.Resource{ // want "resource should configure StateUpgraders instead of MigrateState"
		MigrateState:  migrateStateFunc,
		SchemaVersion: 1,
	}

	// Passing

	_ = schema.Resource{
		SchemaVersion:  1,
		StateUpgraders: []schema.StateUpgrader{},
	}
}

func migrateStateFunc(version int, is *terraform.InstanceState, v interface{}) (*terraform.InstanceState, error) {
	return nil, nil
}
