package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData

	_ = d.HasChange("attr1") || d.HasChange("attr2") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || d.HasChange("attr2") || d.HasChanges("attr3") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"
		d.HasChange("attr2")

	_ = d.HasChanges("attr1", "attr2")

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	_ = d.HasChange("attr1") || d.HasChange("attr2") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || d.HasChange("attr2") || d.HasChanges("attr3") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"
		d.HasChange("attr2")

	_ = d.HasChanges("attr1", "attr2")

	return nil
}
