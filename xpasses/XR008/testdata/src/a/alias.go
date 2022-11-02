package a

import (
	"context"
	e "os/exec"
)

func fAlias() {
	failingAlias := e.CommandContext // want "avoid os/exec.CommandContext"

	e.CommandContext(context.Background(), "true") // want "avoid os/exec.CommandContext"

	failingAlias(context.Background(), "true")
}
