package a

import (
	"testing"
)

func TestAccThing1_import(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_importBasic(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_Import(t *testing.T) {} // want "acceptance test function name should not include import"

func TestAccThing1_basic(t *testing.T) {}

func TestAccThing1_requiresImport(t *testing.T) {}
