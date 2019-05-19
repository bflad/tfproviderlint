package a

import (
	"testing"
)

func TestAccThing1_import(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_importBasic(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_basicImport(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_basic(t *testing.T) {}
