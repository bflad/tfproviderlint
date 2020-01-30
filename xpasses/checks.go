package xpasses

import (
	"github.com/bflad/tfproviderlint/xpasses/XR001"
	"github.com/bflad/tfproviderlint/xpasses/XR002"
	"github.com/bflad/tfproviderlint/xpasses/XR003"
	"github.com/bflad/tfproviderlint/xpasses/XR004"
	"github.com/bflad/tfproviderlint/xpasses/XS001"
	"golang.org/x/tools/go/analysis"
)

// AllChecks contains all Analyzers that report issues
// This can be consumed via multichecker.Main(xpasses.AllChecks...) or by
// combining these Analyzers with additional custom Analyzers
var AllChecks = []*analysis.Analyzer{
	XR001.Analyzer,
	XR002.Analyzer,
	XR003.Analyzer,
	XR004.Analyzer,
	XS001.Analyzer,
}
