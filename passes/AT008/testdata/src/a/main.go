package a

import (
	"testing"
)

// Comment ignored

//lintignore:AT008
func TestAccThing1_commentIgnored1(invalid *testing.T) {}

// Failing

func TestAccThing1_failing1(invalid *testing.T) {} // want "acceptance test function declaration \\*testing.T parameter should be named t"

// Passing

func TestAccThing1_passing1(t *testing.T) {}
