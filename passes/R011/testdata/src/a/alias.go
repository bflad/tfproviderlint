package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Resource{ // want "resource should configure StateUpgraders instead of MigrateState"
		MigrateState:  migrateStateFunc,
		SchemaVersion: 1,
	}
}
