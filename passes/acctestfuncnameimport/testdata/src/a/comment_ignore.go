package a

import (
	"testing"
)

func TestAccThingCommentIgnore1_import(t *testing.T) {} //lintignore:acctestfuncnameimport

func TestAccThingCommentIgnore1_importBasic(t *testing.T) {} //lintignore:acctestfuncnameimport

func TestAccThingCommentIgnore1_basicImport(t *testing.T) {} //lintignore:acctestfuncnameimport

//lintignore:acctestfuncnameimport
func TestAccThingCommentIgnore2_import(t *testing.T) {}

//lintignore:acctestfuncnameimport
func TestAccThingCommentIgnore2_importBasic(t *testing.T) {}

//lintignore:acctestfuncnameimport
func TestAccThingCommentIgnore2_basicImport(t *testing.T) {}
