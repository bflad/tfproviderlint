package a

import (
	e "os/exec"
)

func fAlias() {
	failingAlias := e.Command // want "avoid os/exec.Command"

	e.Command("true") // want "avoid os/exec.Command"

	failingAlias("true")
}
