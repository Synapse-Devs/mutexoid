package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
)

func main() {
	registry := analyzers.NewRegistry()
	singlechecker.Main(registry.Get("testpackage"))
}
