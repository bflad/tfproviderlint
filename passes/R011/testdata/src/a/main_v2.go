package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func f_v2() {
	// Comment ignored

	//lintignore:R011
	_ = schema.Resource{
		MigrateState:  migrateStateFunc_v2,
		SchemaVersion: 1,
	}

	// Failing

	_ = schema.Resource{ // want "resource should configure StateUpgraders instead of MigrateState"
		MigrateState:  migrateStateFunc_v2,
		SchemaVersion: 1,
	}

	// Passing

	_ = schema.Resource{
		SchemaVersion:  1,
		StateUpgraders: []schema.StateUpgrader{},
	}
}

func migrateStateFunc_v2(version int, is *terraform.InstanceState, v interface{}) (*terraform.InstanceState, error) {
	return nil, nil
}
