package a

import (
	"context"
	"os/exec"
)

var failing = exec.CommandContext // want "avoid os/exec.CommandContext"

func f() {
	// Comment ignored

	//lintignore:XR008
	exec.CommandContext(context.Background(), "true")

	exec.CommandContext(context.Background(), "true") //lintignore:XR008

	// Failing

	exec.CommandContext(context.Background(), "true") // want "avoid os/exec.CommandContext"

	failing(context.Background(), "true")
}
