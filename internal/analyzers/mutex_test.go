package analyzers_test

import (
	"testing"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMutexAnalyzer(t *testing.T) {
	analyzer := analyzers.NewRegistry().Get("mutexoid")
	analysistest.Run(t, analysistest.TestData(), analyzer, "mutexcheck")
}
