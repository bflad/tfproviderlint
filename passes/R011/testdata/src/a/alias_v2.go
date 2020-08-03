package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	_ = s.Resource{ // want "resource should configure StateUpgraders instead of MigrateState"
		MigrateState:  migrateStateFunc_v2,
		SchemaVersion: 1,
	}
}
