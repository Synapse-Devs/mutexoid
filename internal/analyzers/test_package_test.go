package analyzers_test

import (
	"testing"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTestPackage(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzers.TestPackageAnalyzer, "testpackage")
}
