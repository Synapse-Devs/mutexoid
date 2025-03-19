package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/Synapse-Devs/mutexoid/internal/analyzers"
)

func main() {
	registry := analyzers.NewRegistry()
	multichecker.Main(registry.GetAll()...)
}
