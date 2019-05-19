package a

import (
	"testing"
)

func TestAccThingCommentIgnore1(t *testing.T) {} //lintignore:acctestfuncnameunderscore

//lintignore:acctestfuncnameunderscore
func TestAccThingCommentIgnore2(t *testing.T) {}
