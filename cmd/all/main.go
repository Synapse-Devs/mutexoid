package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"github.com/Synapse-Devs/mutexoid/internal/analyzers/testcases"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		analyzers.MutexAnalyzer,
		analyzers.EnglishCommentsAnalyzer,
		analyzers.TestPackageAnalyzer,
		analyzers.ParallelTests,
		testcases.Analyzer,
	)
}
