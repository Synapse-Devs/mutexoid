package analyzer_test

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzer"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := filepath.Join("testdata", "src", "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "testdata")
}
