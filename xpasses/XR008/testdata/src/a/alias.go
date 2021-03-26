package a

import (
	"context"
	e "os/exec"
)

var failingAlias = e.CommandContext // want "avoid os/exec.CommandContext"

func fAlias() {
	e.CommandContext(context.Background(), "true") // want "avoid os/exec.CommandContext"

	failingAlias(context.Background(), "true")
}
