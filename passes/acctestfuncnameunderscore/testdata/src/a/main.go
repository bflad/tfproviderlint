package a

import (
	"testing"
)

func TestAccThing1(t *testing.T) {} // want "acceptance test function name should include underscore"

func TestAccThing2_basic(t *testing.T) {}
