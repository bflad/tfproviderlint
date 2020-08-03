package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d s.ResourceData

	_ = d.HasChange("attr1") || d.HasChange("attr2") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || d.HasChange("attr2") || d.HasChange("attr3") // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"

	_ = d.HasChange("attr1") || // want "multiple ResourceData.HasChange\\(\\) calls can be combined with single HasChanges\\(\\) call"
		d.HasChange("attr2")

	_ = d.HasChanges("attr1", "attr2")
}
