package xpasses

import (
	"github.com/bflad/tfproviderlint/xpasses/XS001"
	"golang.org/x/tools/go/analysis"
)

// AllChecks contains all Analyzers that report issues
// This can be consumed via multichecker.Main(xpasses.AllChecks...) or by
// combining these Analyzers with additional custom Analyzers
var AllChecks = []*analysis.Analyzer{
	XS001.Analyzer,
}
