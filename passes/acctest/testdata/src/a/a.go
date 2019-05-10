// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package a

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func f() {
	_ = resource.TestCase{} // want "missing CheckDestroy"

	var t *testing.T                      // want "missing CheckDestroy"
	resource.Test(t, resource.TestCase{}) // want "missing CheckDestroy"
}
