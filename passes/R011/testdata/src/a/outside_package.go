package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Resource{
		MigrateState:  migrateStateFunc_v2,
		SchemaVersion: 1,
	}
}
