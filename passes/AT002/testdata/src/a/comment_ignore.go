package a

import (
	"testing"
)

func TestAccThingCommentIgnore1_import(t *testing.T) {} //lintignore:AT002

func TestAccThingCommentIgnore1_importBasic(t *testing.T) {} //lintignore:AT002

func TestAccThingCommentIgnore1_Import(t *testing.T) {} //lintignore:AT002

//lintignore:AT002
func TestAccThingCommentIgnore2_import(t *testing.T) {}

//lintignore:AT002
func TestAccThingCommentIgnore2_importBasic(t *testing.T) {}

//lintignore:AT002
func TestAccThingCommentIgnore2_Import(t *testing.T) {}
