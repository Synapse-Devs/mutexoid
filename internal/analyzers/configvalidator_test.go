package analyzers_test

import (
	"testing"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestConfigValidatorAnalyzer(t *testing.T) {
	analyzer := analyzers.NewRegistry().Get("configvalidator")
	analysistest.Run(t, analysistest.TestData(), analyzer, "configvalidator")
}
