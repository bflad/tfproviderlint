package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func f() {
	// Passing

	_ = validation.ListOfUniqueStrings

	// Comment ignored

	//lintignore:V006
	_ = validation.ValidateListUniqueStrings

	_ = validation.ValidateListUniqueStrings //lintignore:V006

	// Failing

	_ = validation.ListOfUniqueStrings // want "deprecated validation.ValidateListUniqueStrings should be replaced with validation.ListOfUniqueStrings"
}
