package main

import (
	"github.com/Synapse-Devs/mutexoid/internal/analyzers/testcases"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(testcases.Analyzer)
}
