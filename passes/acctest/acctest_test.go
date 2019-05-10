// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package acctest_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/terraform-providers/terraform-provider-aws/linter/passes/acctest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, acctest.Analyzer, "a")
}
