package a

import (
	e "os/exec"
)

var failingAlias = e.Command // want "avoid os/exec.Command"

func fAlias() {
	e.Command("true") // want "avoid os/exec.Command"

	failingAlias("true")
}
