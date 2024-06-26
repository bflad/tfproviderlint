package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

const (
	ConstKey = "ConstValue"
)

func f() {
	// Comment ignored

	//lintignore:AT009
	_ = acctest.RandStringFromCharSet(1, acctest.CharSetAlpha)

	// Failing

	_ = acctest.RandStringFromCharSet(1, acctest.CharSetAlphaNum) // want "should use RandString call instead"

	// Passing
	_ = acctest.RandStringFromCharSet(1, acctest.CharSetAlpha)
	_ = acctest.RandStringFromCharSet(1, "abc123")
	_ = acctest.RandString(1)
}
