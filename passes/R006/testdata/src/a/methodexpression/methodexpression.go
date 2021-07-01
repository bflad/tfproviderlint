package methodexpression

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type RetryError = resource.RetryError

var RetryableError = resource.RetryableError
