package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		MigrateState:  migrateStateFunc,
		SchemaVersion: 1,
	}
}
