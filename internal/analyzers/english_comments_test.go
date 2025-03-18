package analyzers_test

import (
	"testing"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEnglishComments(t *testing.T) {
	analyzer := analyzers.NewRegistry().Get("englishcomments")
	analysistest.Run(t, analysistest.TestData(), analyzer, "englishcomments")
}
