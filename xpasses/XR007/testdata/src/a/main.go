package a

import (
	"os/exec"
)

func f() {
	failing := exec.Command // want "avoid os/exec.Command"
	// Comment ignored

	//lintignore:XR007
	exec.Command("true")

	exec.Command("true") //lintignore:XR007

	// Failing

	exec.Command("true") // want "avoid os/exec.Command"

	failing("true")
}
