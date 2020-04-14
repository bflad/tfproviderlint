package a

import (
	v "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func falias() {
	// Passing

	_ = v.ListOfUniqueStrings

	// Comment ignored

	//lintignore:V006
	_ = v.ValidateListUniqueStrings

	_ = v.ValidateListUniqueStrings //lintignore:V006

	// Failing

	_ = v.ListOfUniqueStrings // want "deprecated v.ValidateListUniqueStrings should be replaced with v.ListOfUniqueStrings"
}
