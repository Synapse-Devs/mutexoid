package analyzers

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestParallelTests(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, ParallelTests, "parallel_tests")
}
