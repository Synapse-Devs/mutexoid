package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		analyzers.MutexAnalyzer,
		analyzers.EnglishCommentsAnalyzer,
		analyzers.TestPackageAnalyzer,
		analyzers.ParallelTests,
	)
}
